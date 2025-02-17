CREATE TABLE IF NOT EXISTS account
(
    id                              UUID PRIMARY KEY NOT NULL,
    customer_id                     UUID
);