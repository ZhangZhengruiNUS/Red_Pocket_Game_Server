CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL DEFAULT ' ',
  "password" varchar NOT NULL DEFAULT ' ',
  "credit" int NOT NULL DEFAULT 0,
  "coupon" int NOT NULL DEFAULT 0,
  "role_type" int NOT NULL DEFAULT 0,
  "create_time" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "inventories" (
  "user_id" bigint,
  "item_id" bigint,
  "quantity" int NOT NULL DEFAULT 0,
  "create_time" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "item_id")
);

CREATE TABLE "items" (
  "item_id" bigserial PRIMARY KEY,
  "item_name" varchar NOT NULL DEFAULT ' ',
  "describe" varchar NOT NULL DEFAULT ' ',
  "pic_path" varchar NOT NULL DEFAULT ' ',
  "price" int NOT NULL DEFAULT 0,
  "reviser_id" bigint DEFAULT null,
  "revise_time" timestamp DEFAULT null,
  "creator_id" bigint NOT NULL DEFAULT -1,
  "create_time" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "game_difficulty_settings" (
  "diff_lv" int PRIMARY KEY,
  "award_density" int NOT NULL DEFAULT 0,
  "enemy_density" int NOT NULL DEFAULT 0,
  "reviser_id" bigint DEFAULT null,
  "revise_time" timestamp DEFAULT null,
  "creator_id" bigint NOT NULL DEFAULT -1,
  "create_time" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "prizes" (
  "prize_id" bigserial PRIMARY KEY,
  "prize_name" varchar NOT NULL DEFAULT ' ',
  "pic_path" varchar NOT NULL DEFAULT ' ',
  "weight" int NOT NULL DEFAULT 0,
  "reviser_id" bigint DEFAULT null,
  "revise_time" timestamp DEFAULT null,
  "creator_id" bigint NOT NULL DEFAULT -1,
  "create_time" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("user_name");

CREATE INDEX ON "inventories" ("user_id");

COMMENT ON COLUMN "users"."role_type" IS '0-normal user 1-admin';

ALTER TABLE "inventories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "inventories" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("item_id");

ALTER TABLE "items" ADD FOREIGN KEY ("reviser_id") REFERENCES "users" ("user_id");

ALTER TABLE "items" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("user_id");

ALTER TABLE "game_difficulty_settings" ADD FOREIGN KEY ("reviser_id") REFERENCES "users" ("user_id");

ALTER TABLE "game_difficulty_settings" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("user_id");

ALTER TABLE "prizes" ADD FOREIGN KEY ("reviser_id") REFERENCES "users" ("user_id");

ALTER TABLE "prizes" ADD FOREIGN KEY ("creator_id") REFERENCES "users" ("user_id");
