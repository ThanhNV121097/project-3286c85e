CREATE TABLE greetings (
  id smallint PRIMARY KEY CHECK (id = 1),
  text text NOT NULL CHECK (text <> ''),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO greetings (id, text)
VALUES (1, 'Hello Word')
ON CONFLICT (id) DO NOTHING;
