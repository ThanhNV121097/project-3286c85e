# ERD — hello-word-22

## Tables

### greetings

Stores the single greeting value displayed on the page.

| Column | Type | Null | Default | Notes |
|---|---|---|---|---|
| `id` | `smallint` | no | none | Primary key; constrained to `1` so table holds one row |
| `text` | `text` | no | none | Exact greeting text shown to guests |
| `created_at` | `timestamptz` | no | `now()` | Row creation time |
| `updated_at` | `timestamptz` | no | `now()` | Last text update time |

Constraints:

- Primary key: `greetings_pkey (id)`.
- Check: `greetings_single_row CHECK (id = 1)`.
- Check: `greetings_text_not_empty CHECK (length(text) > 0)`.

Seed data:

```sql
INSERT INTO greetings (id, text) VALUES (1, 'Hello Word')
ON CONFLICT (id) DO NOTHING;
```

## Relationships

No relationships. Project has one table and no users, sessions, or audit entities.
