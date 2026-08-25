# Architecture Overview — hello-word-22

## Shape and stack

Fullstack: Next.js frontend, Go backend, PostgreSQL database.

| Part | Choice | Reason | Rejected alternative |
|---|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3 | Matches project default and container contract | Static HTML rejected because text must come from API |
| Backend | Go 1.22 HTTP service | Small API, fast build, existing pipeline default | Node API rejected to avoid second runtime pattern |
| Database | PostgreSQL 16 | Required one stored greeting row | In-memory value rejected because frontend must prove DB path |
| Run | `docker compose up` from repo root | Boots DB, backend, frontend together | Separate commands rejected because pipeline must test end-to-end |

## Repository layout

```text
code/backend/              Go module
code/backend/cmd/api/      single main package and HTTP server
code/backend/internal/     backend packages owned by service
code/backend/migrations/   timestamped .up.sql/.down.sql migration pairs
code/frontend/app/         App Router composition root and global styles
code/frontend/components/  story components, PascalCase filenames
docs/architecture/         project technical contracts
```

`app/page.tsx` stays thin composition root. Story components own markup and mount there with one import and one element.

## Data flow

Browser loads frontend. Frontend story code calls `NEXT_PUBLIC_API_URL + /v1/greeting`. Backend reads one row from PostgreSQL and returns `{ "greeting": "Hello Word" }`. No auth, no write path, no cache layer.

## Environment variables

| Service | Key | Use |
|---|---|---|
| backend | `DATABASE_URL` | PostgreSQL connection string injected by runtime |
| backend | `PORT` | HTTP listen port, default `8080` |
| backend | `APP_PORT` | fallback listen port when `PORT` absent |
| frontend | `NEXT_PUBLIC_API_URL` | browser-visible backend origin |
| root compose | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` | local database bootstrap |
| root compose | `BACKEND_PORT`, `FRONTEND_PORT` | host port overrides |

Every service has `.env.example` with comments only, no secrets.

## Backend conventions

One `main` package only: `cmd/api`. On startup: read `DATABASE_URL`, connect, apply every pending migration from embedded `migrations/`, then listen on `PORT`, fallback `APP_PORT`, fallback `8080`. `/healthz` returns 200 only after migrations succeeded and `SELECT 1` succeeds.

Routes use `/v1/...`; no `/api` prefix. Errors use shared envelope from `docs/architecture/services.md`.

## Frontend conventions

Server Components by default. Files using hooks, event handlers, or browser APIs must start with literal first line `"use client"`. Components use `export default function ComponentName()`. Shared visual values live in `app/globals.css` tokens only: colour, spacing, typography, radius, shadow, motion.

## CI and quality gate

`.github/workflows/ci.yml` runs on pull requests: `go build ./...`, `go vet ./...`, `go test ./...`, `npm ci`, `npm run lint`, `npm run build`, `npm test --if-present`, and CSS token checks. Docker/container files are precommitted pipeline contracts; do not edit workflows unless assigned.

## Run locally

1. Copy `.env.example` to `.env` at repo root if overriding defaults.
2. Run `docker compose --profile local up --build` from repo root.
3. Open frontend at `http://localhost:3000`; backend health at `http://localhost:8080/healthz`.

## Known risks

| Risk | Mitigation |
|---|---|
| Missing seed row would make story fail | Initial migration inserts one `Hello Word` row |
| Backend healthy before DB ready | Health endpoint performs `SELECT 1` |
| Hardcoded visual values drift from design | CI token checks and minimal token set in `globals.css` |
