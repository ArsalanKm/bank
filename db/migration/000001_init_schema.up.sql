CREATE TABLE "accounts" (
"id" bigserial PRIMARY KEY,
"owner" VARCHAR NOT NULL,
"balance" bigint NOT NULL,
"currency" VARCHAR NOT null,
"created_at" TIMESTAMP NOT NULL DEFAULT (now())
);


CREATE TABLE "entries" (
    "id" bigserial PRIMARY KEY,
    "account_id" bigint,
    "amount" bigint NOT null,
    "created_at" TIMESTAMP not NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
    "id" bigserial PRIMARY KEY,
    "from_account_id" bigint,
    "to_account_id" bigint,
    "amount" bigint NOT null,
    "created_at" TIMESTAMP not NULL DEFAULT (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");


