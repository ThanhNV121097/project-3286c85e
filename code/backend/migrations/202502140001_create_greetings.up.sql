CREATE TABLE IF NOT EXISTS schema_migrations (
  version text PRIMARY KEY,
  applied_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS greetings (
  id smallint PRIMARY KEY,
  text text NOT NULL CHECK (length(text) > 0),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  CONSTRAINT greetings_single_row CHECK (id = 1)
);

INSERT INTO greetings (id, text)
VALUES (1, 'Hello Word')
ON CONFLICT (id) DO NOTHING;
