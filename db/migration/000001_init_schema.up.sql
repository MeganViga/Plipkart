CREATE TABLE "user_data" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "role" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "created_at" timestamptz NOT NULL,
  "modified_at" timestamptz NOT NULL
);

CREATE TABLE "user_address" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "address_line1" varchar NOT NULL,
  "address_line2" varchar NOT NULL,
  "city" varchar NOT NULL,
  "postal_code" varchar NOT NULL,
  "country" varchar NOT NULL,
  "phone_number" varchar NOT NULL,
  "created_at" timestamptz NOT NULL,
  "modified_at" timestamptz NOT NULL
);

CREATE TABLE "user_payment" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "card_no" varchar NOT NULL,
  "card_user_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL,
  "modified_at" timestamptz NOT NULL
);

CREATE TABLE "product" (
  "id" BIGSERIAL PRIMARY KEY,
  "product_name" varchar NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL,
  "modified_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

CREATE TABLE "product_category" (
  "product_id" bigserial NOT NULL,
  "color" varchar NOT NULL,
  "quantity" int NOT NULL
);

CREATE TABLE "order" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "total_price" bigint NOT NULL,
  "status" varchar NOT NULL
);

CREATE TABLE "order_details" (
  "order_id" bigserial NOT NULL,
  "product_id" bigserial NOT NULL,
  "color" varchar NOT NULL,
  "quantity" int NOT NULL,
  "price" bigint NOT NULL
);

CREATE UNIQUE INDEX ON "product_category" ("product_id", "color");

CREATE UNIQUE INDEX ON "order" ("id", "user_id");

CREATE UNIQUE INDEX ON "order_details" ("order_id", "product_id");

ALTER TABLE "user_address" ADD FOREIGN KEY ("user_id") REFERENCES "user_data" ("id");

ALTER TABLE "user_payment" ADD FOREIGN KEY ("user_id") REFERENCES "user_data" ("id");

ALTER TABLE "product_category" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("user_id") REFERENCES "user_data" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");
