CREATE TABLE IF NOT EXISTS "event"
(
    id                              TEXT PRIMARY KEY NOT NULL,
    event_type                      TEXT,
    entity_id                       TEXT,
    created_at                      TIMESTAMP DEFAULT now(),
    status							CHAR DEFAULT 'C'
);