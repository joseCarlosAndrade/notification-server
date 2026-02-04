# Notification Server

Simple notification server implementation using redis and redpanda.

maybe use some sort of db to store it?

mongodb

## useful notes

you can subscribe to different channels at once by:

```go
pubsub := s.rdb.PSubscriber(context.Background(), "tasks.*")
```

## main goal

- implement my own log wrapper from uber zap
- implement my own otel wrapper from otel

have a way of creating notifications:

- publishing it on a redpanda topic(s) (any external service could do it asynchronously)
- have this notification service listen to this topic(s) to gather all notifications from all services
- make it store on some sort of db (mongodb?)
- expose http endpoint that
  - reads all notification for that service
  - it must be able to receives read confirmation for a specific one
  - db must mark notifications that are already read
  - fetch must be able to fetch both: non-read and all notifications
  - use redis to cache notifications to avoid duplicate reads
- this service must be scalable (making different replicas dont affect its funcionality, but it distributes the work)
- use load balancer to balace requests (http)
- use consumer groups to differ new events
- use health probes

## local compose config

### redpanda

```text
APP_REDPANDABROKERS="localhost:19092"
APP_KAFKACONSUMERGROUP="notification-group"
APP_NOTIFICATIONTOPIC="notification.events.v1"
```

### mongo

```text
mongodb://admin:password@localhost:27017
```

database name: `notifications`
colletion name: `notifications`
