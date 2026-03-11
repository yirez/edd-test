# Event-Driven Demo Environment

## Goal

Build a lightweight Go monorepo to learn event-driven architecture patterns through code, not framework complexity.

Primary patterns:

- transactional outbox
- idempotent consumer / inbox
- retries and dead-letter queue
- choreography saga
- orchestration saga
- compensation
- CQRS projection
- correlation and causation IDs

## Roadmap

Use an `order` flow as the demo domain.

Expected flow:

1. Create order over HTTP.
2. Persist order and publish `order.created` through the outbox.
3. Reserve inventory from the inventory service.
4. Authorize payment from the payment service.
5. Update order state from downstream events.
6. Notify on final outcome.
7. Build read models from the event stream.

## Service Shape

Per service, prefer:

- `internal/globals`
- `internal/usecase`
- `internal/repo`
- `internal/types`
- `config.yml`
- top-level `router.go` when the service exposes HTTP

Notes:

- `config.yml` should be the common per-app configuration basis.
- `example_config.yml` should be committed per app so developers have a safe starting point.
- router wiring can live in `cmd/<service>/router.go` to keep HTTP entrypoints shallow.
- use case files should group service operations behind one app-level `UseCase`.
- repository interfaces should prefer `FooRepo` naming and stay narrow.
- usecases should prefer `FooUseCase` plus `FooUseCaseImpl` with a `NewFooUseCase(...)` constructor.
- inject only the repo or setting a usecase actually needs; do not pass broad globals when they are unused.
- `app/` is intentionally omitted.

## Shared Code

Current shared code is intentionally small:

- `pkg/common_types` for types and contracts shared across services
- shared `github.com/yirez/go-gw-test/pkg/configuration_manager` for config, logger, and DB bootstrap
- shared `github.com/yirez/go-gw-test/pkg/rest_qol` for operational routes, request IDs, access logging, and HTTP server startup

Keep adding shared packages only when they are actually reused.

## Current Direction

Scaffolded now:

- `order-service`
- `inventory-service`
- `projection-service`
- `notification-service`
- compose setup for Postgres and Kafka
- initial shared contracts and types under `pkg/common_types`

Add next:

- order persistence with GORM/Postgres
- Kafka producer and consumer wiring
- transactional outbox
- idempotent consumer / inbox
- `payment-service`
- compensation flow
- choreography saga

Add later:

- `saga-service` for orchestration comparison

## Documentation Rule

Keep this file as a roadmap. Add details as code is implemented. Do not predeclare packages or apps that are not yet needed.
