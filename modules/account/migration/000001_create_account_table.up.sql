CREATE TABLE IF NOT EXISTS account
(
    id                              TEXT PRIMARY KEY NOT NULL,
    customer_id                     TEXT,
    created_at                      TIMESTAMP default now(),
    status							char default 'C'
);