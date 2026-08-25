# Architecture Overview — hello-word-22

## Shape and stack

Fullstack project: Next.js frontend, Go backend, PostgreSQL database.

| Part | Choice | Reason |
|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3 | Matches repo scaffold and one-page UI need |
| Backend | Go 1.22 HTTP API | Small compiled service, matches default stack |
| Database | PostgreSQL 16 | Required because greeting lives in one stored row |
| Runtime | `docker compose up` from repo root | Boots database, backend, frontend together |

## Folder layout

```text
code/backend/                  Go API service
  cmd/api/main.go              process entrypoint
  internal/migrations/         embedded SQL migration runner
  migrations/                  timestamped .up.sql/.down.sql files
code/frontend/                 Next.js app
  app/layout.tsx               root document shell
  app/page.tsx                 composition root only
  app/globals.css              shared tokens and base styles
  components/                  story components later
  lib/                         frontend helpers later
docs/architecture/overview.md  this file
docs/architecture/erd.md       table design
docs/architecture/services.md  API contracts
```

## Data flow

1. Browser loads frontend.
2. Frontend story component will call `GET /v1/greeting` through `NEXT_PUBLIC_API_URL`.
3. Backend reads one active row from PostgreSQL.
4. Frontend renders returned `greeting` exactly as JSON value.

## Service boundaries

Frontend owns presentation only. It must not hardcode `Hello Word` outside tests or mocks. Backend owns validation, database access, migrations, and response envelope. Database owns durable greeting row and seed data.

## Env vars

Backend reads:

```text
DATABASE_URL   PostgreSQL connection string injected by runtime
PORT           listen port, default 8080
APP_PORT       fallback listen port when PORT is unset
```

Frontend reads:

```text
NEXT_PUBLIC_API_URL   browser-visible backend base URL
```

Root compose reads PostgreSQL credentials plus optional port and image settings from `.env`.

## Naming conventions

- Backend routes use `/v1/...`; never `/api/...` because deploy proxy strips `/api` before Go receives request.
- Backend JSON errors use shared envelope from `services.md`.
- SQL migration filenames sort by timestamp and end in `.up.sql` / `.down.sql`.
- React component files use PascalCase and `export default function ComponentName()`.
- `app/page.tsx` stays server component and only composes child components.
- CSS modules use tokens from `app/globals.css`; no hardcoded colors, spacing, or token fallbacks.

## Decisions

| Decision | Rejected alternative | Tradeoff |
|---|---|---|
| Self-migrate on backend boot | Manual migration step | Slower startup, but empty runtime DB becomes usable automatically |
| One `greetings` row seeded by migration | Hardcoded backend constant | Adds DB read, but proves stored data pipeline |
| `/healthz` checks DB after migrations | Process-only health | Health can fail during DB outage, but avoids false healthy app |
| Plain `net/http` server | Router dependency | Fewer features, but no extra package for two endpoints |
| Frontend scaffold renders neutral shell only | Implement finished greeting page now | Keeps scaffold task clean; story owns UI implementation |

## Security and reliability

- Backend exits if `DATABASE_URL` is missing.
- SQL uses parameterized queries in story code; scaffold migration runner executes trusted embedded SQL only.
- Health returns 200 only after migrations and `SELECT 1` succeed.
- Public page has no auth, no cookies, no user input.

## Run and verify

```bash
cp .env.example .env
docker compose --profile local up --build
```

CI gate in `.github/workflows/ci.yml` runs:

- `go build ./...`, `go vet ./...`, `go test ./...`
- `npm ci`, `npm run lint`, `npm run build`, `npm test --if-present`
- CSS token checks for modules and globals usage

## Known limits

- No edit API; greeting value changes only by database update or migration.
- No frontend loading/error UI because approved design has default state only.
- No observability stack; container logs are enough for one endpoint.
