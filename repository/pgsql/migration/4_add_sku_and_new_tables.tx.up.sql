-- 4_add_sku_and_new_tables.tx.up.sql

-- Добавление нового поля sku в таблицу ad_campaigns
ALTER TABLE "ad_campaigns"
    ADD COLUMN "sku" BIGINT;

-- Создание таблицы bidder_requests
CREATE TABLE IF NOT EXISTS "bidder_requests" (
                                                 "id" INTEGER PRIMARY KEY,
                                                 "request" TEXT
);

-- Создание таблицы bidder_lists
CREATE TABLE IF NOT EXISTS "positions" (
                                           "request_id" INTEGER REFERENCES "bidder_requests" ("id") ON DELETE CASCADE,
                                           "sku" BIGINT,
                                           "organic" INTEGER,
                                           "position" INTEGER,
                                           "updated_at" TIMESTAMPTZ DEFAULT now(),
                                           PRIMARY KEY ("request_id")
);

-- Создание таблицы positions
CREATE TABLE IF NOT EXISTS "bidder_lists" (
                                              "ad_id" BIGINT REFERENCES "ad_campaigns" ("ad_id") ON DELETE CASCADE,
                                              "request_id" INTEGER REFERENCES "bidder_requests" ("id"),
                                              "current_bid" INTEGER,
                                              "max_bid" INTEGER,
                                              "paused" BOOLEAN DEFAULT FALSE,
                                              "created_at" TIMESTAMPTZ DEFAULT now(),
                                              "updated_at" TIMESTAMPTZ DEFAULT now(),
                                              PRIMARY KEY ("ad_id")
);