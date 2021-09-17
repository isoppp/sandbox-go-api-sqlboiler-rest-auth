CREATE TABLE "users"
(
    "id"              SERIAL PRIMARY KEY,
    "email"           varchar UNIQUE NOT NULL,
    "hashed_password" varchar        NOT NULL,
    "created_at"      timestamptz     NOT NULL DEFAULT (now()),
    "updated_at"      timestamptz     NOT NULL DEFAULT (now())
);

CREATE TABLE "roles"
(
    "id"         SERIAL PRIMARY KEY,
    "name"       varchar UNIQUE NOT NULL,
    "created_at" timestamptz     NOT NULL DEFAULT (now()),
    "updated_at" timestamptz     NOT NULL DEFAULT (now())
);

CREATE TABLE "user_role"
(
    "id"      SERIAL PRIMARY KEY,
    "user_id" int NOT NULL,
    "role_id" int NOT NULL
);

CREATE TABLE "posts"
(
    "id"         SERIAL PRIMARY KEY,
    "content"    text       NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "user_id"    int        NOT NULL
);

ALTER TABLE "user_role"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_role"
    ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");

ALTER TABLE "posts"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE UNIQUE INDEX ON "user_role" ("user_id", "role_id");

CREATE INDEX ON "posts" ("created_at");

CREATE INDEX ON "posts" ("updated_at");
