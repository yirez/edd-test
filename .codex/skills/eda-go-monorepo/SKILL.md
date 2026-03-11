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
- `internal/usecase` for behavior-oriented application logic, including event handlers and workers
- `internal/repo` for persistence adapters and repository interfaces local to the service
- `internal/types` for service-local types that should not be shared
- `config.yml` at the service root for per-app configuration
- `router.go` at the service root when the service exposes HTTP

Do not add `internal/app`.

Prefer use case-oriented files over extra technical layers. Logic that might otherwise live under `domain`, `publisher`, or `consumer` should usually live under `internal/usecase`.

## Shared Code

- Use `pkg/` only for code reused across services.
- Prefer `pkg/common_types` for shared contracts and shared common types.
- Prefer shared `github.com/yirez/go-gw-test/pkg/configuration_manager` for config, logger, and DB bootstrap.
- Prefer shared `github.com/yirez/go-gw-test/pkg/rest_qol` for HTTP/server helpers.
- Keep shared contracts explicit and small.
- Put service-local types in `internal/types`.
- Put intentionally shared types and contracts under `pkg/common_types`.

## Repositories

- Prefer `FooRepo` naming for repository interfaces.
- Prefer `FooRepoImpl` for the concrete implementation.
- Prefer `NewFooRepo()` constructors returning the interface.
- Keep repository interfaces small and use case-oriented.
- Add concrete repo implementations only when the implementation work starts to exist.

## Usecases

- Prefer `FooUseCase` naming for usecase interfaces.
- Prefer `FooUseCaseImpl` for the concrete implementation.
- Prefer `NewFooUseCase(...)` constructors returning the interface.
- Prefer one app-level usecase per service that exposes multiple service operations.
- Prefer service methods such as `CreateOrder(...)`, `GetOrderView(...)`, or `HandleEvent(...)` over one usecase type per handler.
- Inject only the specific repo or setting a usecase actually needs.
- Do not pass broad globals into a usecase unless the usecase genuinely uses them.
- Wire constructed usecases in `router.go`, then pass bound methods to HTTP routes.

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
