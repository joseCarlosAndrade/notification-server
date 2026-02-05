package redpanda

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/joseCarlosAndrade/notification-server/internal/core/domain/logger"
	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/models"
	"github.com/joseCarlosAndrade/notification-server/internal/core/domain/port"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.uber.org/zap"
)

// EventsHub implements the EventsHub interface
type EventsHub struct {
	service       *port.Service
	client        *kgo.Client
	topic         string
	consumerGroup string
}

// makes sure EventsHub implements the interface
var _ port.EventsHub = (*EventsHub)(nil)

func NewEventsHub(ctx context.Context, serviceRepository *port.Service, brokers []string, topic, group string) (EventsHub, error) {
	
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(group),
		kgo.ConsumeTopics(topic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)

	if err != nil {
		return EventsHub{}, err
	}

	log.L(ctx).Info("redpanda started and connected",
		zap.String("brokers", brokers[0]),
		zap.String("topic", topic),
		zap.String("group", group))

	return EventsHub{
		client: client,
		topic: topic,
		consumerGroup: group,
		service: serviceRepository,
	}, nil
}

// Run implements port.Runner interface
func (e *EventsHub) Run(ctx context.Context) error {
	for {
		// avoid fetching with canceled context
		select {
		case <-ctx.Done():
			log.L(ctx).Warn("context canceled")
			return nil
		default:
		}
		// log.L(ctx).Info("polling")

		fetches := e.client.PollFetches(ctx)
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()

			err := e.processRecord(ctx, record)
			if err != nil {
				log.L(ctx).Error("error processing record", zap.Error(err))

				// TODO: handle this error: maybe dead letter?
			}
		}
	}
}

// Close implements port.Runner interface
func (e *EventsHub) Close(ctx context.Context) error {
	// todo: close connection and clean up
	e.client.Close()
	return nil
}

// proccesRecord starts a trace_id for this context and process the record to store this entry in
// database and in cache
func (e *EventsHub) processRecord(ctx context.Context, record *kgo.Record) error {
	// loads a new trace_id into the context
	ctx = log.InitResources(ctx)

	log.L(ctx).Debug("processing record", 
		zap.String("key", string(record.Key)), 
		zap.String("value", string(record.Value)))
	
		// NEXT STEPS: VALIDATE THIS AND THE restart: unless-stopped
	notification, err := validatePayload(record.Key, record.Value)
	if err != nil {
		return fmt.Errorf("could not process record: %w", err)
	}

	// save in cache
	// save in db

	err = (*e.service).SaveNewNotification(ctx, notification)
	if err != nil {
		log.L(ctx).Error("could not save new notification", zap.Error(err))
		return fmt.Errorf("error saving notification: %w", err)
	}

	return nil
}

func validatePayload(_ []byte, value []byte) (*models.NotificationRecord, error) {

	var payload models.NotificationRecord

	if err := json.Unmarshal(value, &payload); err != nil {
		return nil, fmt.Errorf("could not parse payload: %w", err)
	}

	// ensuring the timestamp is utc
	if payload.SentAt.Location() != time.UTC {
		payload.SentAt = payload.SentAt.UTC()
	}

	return &payload, nil
}


func (e *EventsHub)IsHealthy(ctx context.Context) error {
	return e.client.Ping(ctx)
}