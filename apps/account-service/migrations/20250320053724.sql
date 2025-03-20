-- Create "account_entities" table
CREATE TABLE "account_entities" (
  "id" text NOT NULL,
  "user_id" text NULL,
  "balance" numeric(20,6) NULL,
  PRIMARY KEY ("id")
);
-- Create "user_entities" table
CREATE TABLE "user_entities" (
  "id" text NOT NULL,
  "name" text NULL,
  "nik" text NULL,
  "phone_number" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_user_entities_nik" to table: "user_entities"
CREATE UNIQUE INDEX "idx_user_entities_nik" ON "user_entities" ("nik");
-- Create index "idx_user_entities_phone_number" to table: "user_entities"
CREATE UNIQUE INDEX "idx_user_entities_phone_number" ON "user_entities" ("phone_number");
