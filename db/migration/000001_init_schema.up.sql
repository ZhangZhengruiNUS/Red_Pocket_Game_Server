CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "role_type" int NOT NULL,
  "create_time" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "inventories" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "item_id" bigint NOT NULL,
  "create_time" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "items" (
  "id" bigserial PRIMARY KEY,
  "item_name" varchar NOT NULL,
  "item_price" int NOT NULL,
  "create_time" timestamp NOT NULL DEFAULT (now()),
  "creator_id" bigint NOT NULL
);

CREATE INDEX ON "accounts" ("username");

CREATE INDEX ON "inventories" ("account_id");

CREATE INDEX ON "items" ("id");

COMMENT ON COLUMN "accounts"."role_type" IS '0-normal user 1-admin';

ALTER TABLE "inventories" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "inventories" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("creator_id") REFERENCES "accounts" ("id");
