CREATE TABLE "sessions"
(
    "id"         int PRIMARY KEY,
    "user_id"    int         NOT NULL,
    "expires_at" timestamp   NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
CREATE INDEX ON "sessions" ("expires_at");
