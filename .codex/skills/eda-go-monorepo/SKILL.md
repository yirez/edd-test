---
name: eda-go-monorepo
description: Lightweight coding conventions for this Go event-driven demo monorepo. Use when creating or changing services, shared packages, compose setup, contracts, or documentation in this repository so structure, pattern visibility, and docs stay consistent.
---

# EDA Go Monorepo

Follow these repo conventions when working in this repository.

## Priorities

- Keep the code lightweight and explicit.
- Optimize for learning event-driven architecture patterns.
- Avoid inventing shared packages or services before they are needed.
- Keep documentation short and current.

## Architecture Docs

- Treat `ARCHITECTURE.md` as a roadmap, not a full inventory.
- Do not list packages, apps, or folders that are only hypothetical.
- Add to `ARCHITECTURE.md` when a component is planned next or already exists.
- Keep `README.md` short and developer-facing.
- Update `ARCHITECTURE.md` and `README.md` in the same change when structure, setup, workflow, or implemented patterns change.

## Service Layout

Use this internal layout per service as needed:

- `internal/globals` for env, config, logger, db, kafka, and service-wide wiring
- `internal/transport` for HTTP handlers and router wiring when the service exposes HTTP
- `internal/usecase` for behavior-oriented application logic, including event handlers and workers
- `internal/repo` for persistence adapters and repository interfaces local to the service
- `internal/types` for service-local types that should not be shared

Do not add `internal/app`.

Prefer use case-oriented files over extra technical layers. Logic that might otherwise live under `domain`, `publisher`, or `consumer` should usually live under `internal/usecase`.

## Shared Code

- Use `pkg/` only for code reused across services.
- Keep shared contracts explicit and small.
- Put service-local types in `internal/types`.
- Put intentionally shared types in `pkg/types`.
- Put shared event and command contracts in `pkg/contracts` when they become necessary.

## Stack

- Use `gorilla/mux` for HTTP routing.
- Use GORM with Postgres for persistence.
- Use Kafka for asynchronous messaging.
- Use one root `go.mod` for the monorepo.

## Pattern Visibility

This repo is for learning, so pattern usage must be obvious in code.

- Add a short local comment where each EDA pattern is implemented.
- Use the form `Pattern: <name> - <why here>.`

Examples:

- `// Pattern: Transactional Outbox - persist event with order update in one DB tx.`
- `// Pattern: Idempotent Consumer - ignore duplicate Kafka message IDs.`
- `// Pattern: Saga Compensation - release inventory after payment failure.`

## Delivery Style

- Start with the smallest useful slice.
- Prefer the order flow as the main demo path.
- Build patterns incrementally: outbox, inbox, retries, projection, saga, compensation.
- Add orchestration only after choreography is clear.
- Keep failure toggles easy to trigger for demos.
