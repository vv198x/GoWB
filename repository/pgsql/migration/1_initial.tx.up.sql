-- 1_initial.tx.up.sql

-- Создание таблицы ad_campaigns
CREATE TABLE IF NOT EXISTS "ad_campaigns" (
                                              "ad_id" BIGINT PRIMARY KEY,
                                              "name" TEXT,
                                              "budget" DOUBLE PRECISION,
                                              "status" INTEGER,
                                              "type" INTEGER,
                                              "created_at" TIMESTAMPTZ DEFAULT now(),
                                              "updated_at" TIMESTAMPTZ DEFAULT now()
);

-- Добавление индексов
CREATE INDEX IF NOT EXISTS idx_ad_campaigns_status ON "ad_campaigns" ("status");
CREATE INDEX IF NOT EXISTS idx_ad_campaigns_type ON "ad_campaigns" ("type");