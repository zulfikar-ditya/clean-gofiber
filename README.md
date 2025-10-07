# Clean Golang fiber

- [ ] CMD
- [ ] API
- [ ] DATABSE
- [ ] DATABSE
- [ ] Toolkit/Utils (Encrypt, Hash, Str, Number, Cache, Etc)
- [ ] Auth
- [ ] ...

#

```
your-service/
├── cmd/
│   ├── api/
│   │   └── main.go          # HTTP entrypoint (Fiber or Gin)
│   └── worker/
│       └── main.go          # background jobs, consumers, schedulers
│
├── domain/                  # business logic (pure Go, no framework)
│   └── order/
│       ├── order.go         # entity / aggregate root
│       ├── item.go          # value object
│       └── service.go       # domain service / rules
│
├── application/             # use-cases, input/output ports
│   └── order/
│       ├── create_order.go  # use-case (CreateOrder)
│       └── ports.go         # interfaces (OrderRepo, EventPublisher)
│
├── adapters/                # interface adapters (HTTP, gRPC, CLI, etc.)
│   └── http/
│       ├── orders_handler.go
│       └── router.go
│
├── infrastructure/          # external implementations
│   ├── db/
│   │   ├── sqlc.yaml
│   │   ├── queries/
│   │   │   └── order.sql
│   │   └── repo_pg.go
│   ├── redis/
│   │   └── cache.go
│   ├── clickhouse/
│   │   └── writer.go
│   ├── kafka/
│   │   └── producer.go      # optional
│   └── logger/
│       └── logger.go
│
├── config/
│   ├── config.go
│   └── env.sample
│
├── events/
│   ├── topics.go
│   └── schema.go
│
├── observability/
│   ├── tracing.go
│   ├── metrics.go
│   └── middleware.go
│
├── migrations/
│   ├── postgres/
│   │   └── 001_init.sql
│   └── clickhouse/
│       └── 001_events.sql
│
├── docs/
│   ├── ADR-001-initial-structure.md
│   └── openapi.yaml
│
├── go.mod
└── Makefile

```
