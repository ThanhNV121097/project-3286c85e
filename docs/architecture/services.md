# Service Contracts — hello-word-22

Base path: `/v1`. Do not mount routes under `/api`; deploy proxy strips that prefix before backend receives request.

## Error envelope

All non-2xx responses use this JSON shape:

```json
{
  "error": {
    "code": "internal_error",
    "message": "Internal server error"
  }
}
```

Rules:

- `code` is stable snake_case for clients and tests.
- `message` is safe for guests, no internals.
- Backend logs internal detail; response stays generic.

## Endpoints

### GET `/v1/greeting`

Reads the stored greeting text.

Request body: none.

Success response: `200 OK`

```json
{
  "greeting": "Hello Word"
}
```

Errors:

| Status | Code | When |
|---|---|---|
| `500` | `internal_error` | Database unavailable or greeting row cannot be read |

### GET `/healthz`

Runtime health check. Not part of public product API.

Request body: none.

Success response: `200 OK`

```json
{
  "status": "ok"
}
```

Errors:

| Status | Code | When |
|---|---|---|
| `503` | `service_unavailable` | Migrations not complete or database `SELECT 1` fails |
