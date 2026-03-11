# edd-test

Lightweight Go monorepo for learning event-driven architecture with Kafka, Postgres, `gorilla/mux`, and GORM.

See [ARCHITECTURE.md](/g:/gitrepo/edd-test/ARCHITECTURE.md) for the current roadmap.

## Repo Rules

- Keep `ARCHITECTURE.md` and `README.md` updated when structure, setup, workflow, or pattern implementation changes.
- Keep docs light and expand them only as the codebase grows.
- Comment each EDA pattern where it is implemented in code.
- Prefer one app-level `UseCase` per service rather than one type per handler.
- Use `config.yml` per app as the common configuration basis.
- Commit `example_config.yml` per app and keep local `config.yml` uncommitted.
- Inject only the specific dependency a usecase needs.

## Current Scaffold

- `order-service`
- `inventory-service`
- `projection-service`
- `notification-service`
- `compose/docker-compose.yml` for Postgres and Kafka
- shared types and contracts under `pkg/common_types`
- shared bootstrap from `github.com/yirez/go-gw-test/pkg/configuration_manager`
- shared HTTP/server helpers from `github.com/yirez/go-gw-test/pkg/rest_qol`

Example comments:

- `// Pattern: Transactional Outbox - persist event with aggregate update in one tx.`
- `// Pattern: Idempotent Consumer - ignore duplicate message ID.`
- `// Pattern: Saga Compensation - release reservation after payment failure.`
