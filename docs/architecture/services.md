# Services — hello-word-22

## Conventions

- Backend routes use `/v1/...` only. Do not mount under `/api`.
- JSON responses use `application/json`.
- Public API; no authentication.
- Error envelope is stable for all endpoints:

```json
{
  "error": {
    "code": "internal_error",
    "message": "Internal server error"
  }
}
```

## Endpoints

### `GET /healthz`

Health endpoint for runtime.

Request: no body.

Success response: `200 OK`

```json
{
  "status": "ok"
}
```

Failure response: `503 Service Unavailable`

```json
{
  "error": {
    "code": "service_unavailable",
    "message": "Service unavailable"
  }
}
```

Rules:

- Return 200 only after migrations succeeded and `SELECT 1` works.
- Return 503 when database check fails.

### `GET /v1/greeting`

Returns stored greeting for page.

Request: no body.

Success response: `200 OK`

```json
{
  "greeting": "Hello Word"
}
```

Failure responses:

| Status | Code | Message | When |
|---|---|---|---|
| 404 | `greeting_not_found` | `Greeting not found` | Required row missing |
| 500 | `internal_error` | `Internal server error` | Database or unexpected failure |

Notes:

- Response object has greeting text only plus no metadata.
- Text must match database value exactly.
