# edd-test

Lightweight Go monorepo for learning event-driven architecture with Kafka, Postgres, `gorilla/mux`, and GORM.

See [ARCHITECTURE.md](/g:/gitrepo/edd-test/ARCHITECTURE.md) for the current roadmap.

## Repo Rules

- Keep `ARCHITECTURE.md` and `README.md` updated when structure, setup, workflow, or pattern implementation changes.
- Keep docs light and expand them only as the codebase grows.
- Comment each EDA pattern where it is implemented in code.

Example comments:

- `// Pattern: Transactional Outbox - persist event with aggregate update in one tx.`
- `// Pattern: Idempotent Consumer - ignore duplicate message ID.`
- `// Pattern: Saga Compensation - release reservation after payment failure.`
