-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "number" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "name" varchar NOT NULL,
  "last_name" varchar,
  "balance" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be positive and negative';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS entries:
-- +goose StatementEnd
