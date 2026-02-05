# Notification Server

Simple notification server implementation using redis and redpanda.

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

## Todos

- implement read operations for mongo package
  - read single notification
  - read all notifications from service
- implement redis caching
- implement api layer
- wrap panic calls so it doesnt quit the app

## Persistence

Using mongo db to store notifications, as well as its status of reading.
Default conn string for local tests:

```text
mongodb://admin:password@localhost:27017
```

## Message schema

Default topic: `notification.events.v1`

```json
{
    "service" : "payments",
    "message" : "new order generated",
    "sentAt" : "2026-02-04T21:34:32Z"
}
```

### Aggregation used to sort last notifications from X time

```bson
[
{
 $match: {
   service : "payments",
   sentAt : {
     $gte :   ISODate('2026-02-04T21:38:32Z')
   }
 } 
},
 {
   $sort: {
     sentAt: -1
   }
 }
]
```

finding past notifications

```golang
// "10 minutes ago"
tenMinsAgo := time.Now().UTC().Add(-10 * time.Minute)
```

## Timestamps

All services must publish records using UTC timezones. This is the core standart to avoid confusion and incoherence betwen services. The timestamps are also stored in UTC.
Universal at the core, Local at the edges.

The edge services should handle the logic to convert the timestamps to its local timezone.

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
