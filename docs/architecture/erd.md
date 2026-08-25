# ERD — hello-word-22

## Tables

### `greetings`

Stores single greeting value shown on page.

| Column | Type | Constraints | Notes |
|---|---|---|---|
| `id` | `smallint` | primary key, `id = 1` | Enforces one-row model |
| `text` | `text` | not null, not empty | Value returned to frontend |
| `created_at` | `timestamptz` | not null, default `now()` | Audit creation time |
| `updated_at` | `timestamptz` | not null, default `now()` | Audit update time |

## Relationships

No relationships. Project has one table and one row.

## Seed data

Migration inserts row:

```text
id: 1
text: Hello Word
```

## Notes

- Empty text blocked by `CHECK (text <> '')`; approved design has no empty state.
- One-row rule blocked by `CHECK (id = 1)` plus primary key.
