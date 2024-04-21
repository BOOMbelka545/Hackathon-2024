-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "number" varchar NOT NULL,
  "password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "name" varchar NOT NULL,
  "last_name" varchar,
  "balance" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "payments" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "place" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("number");

CREATE INDEX ON "payments" ("account_id");

COMMENT ON COLUMN "payments"."amount" IS 'can be positive and negative';

ALTER TABLE "payments" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS payments:
-- +goose StatementEnd
