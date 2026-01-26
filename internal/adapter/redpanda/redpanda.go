package redpanda

import (
	"context"

	log "github.com/joseCarlosAndrade/notification-server/internal/core/domain/logger"
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
	
		// NEXT STEPS: VALIDATE THIS AND THE WAY THIS GROUP CONSUMES, OFFSET AND PARTITION

	// validate payload

	// save on database

	// save in cache

	return nil
}

