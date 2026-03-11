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
- `internal/transport`
- `internal/usecase`
- `internal/repo`
- `internal/types`

Notes:

- `transport` is optional for services with little or no HTTP surface.
- use case files should hold business flow handlers, event handlers, and workers.
- `app/` is intentionally omitted.

## Current Direction

Start with:

- `order-service`
- `inventory-service`
- `projection-service`
- `notification-service`
- compose setup for Postgres and Kafka

Add next:

- `payment-service`
- compensation flow
- choreography saga

Add later:

- `saga-service` for orchestration comparison

## Documentation Rule

Keep this file as a roadmap. Add details as code is implemented. Do not predeclare packages or apps that are not yet needed.
