-- 2_add_history_table.up.sql

-- Создание таблицы history
CREATE TABLE IF NOT EXISTS "histories" (
                                         "id" SERIAL PRIMARY KEY,
                                         "ad_id" BIGINT NOT NULL,
                                         "date" DATE NOT NULL,
                                         "amount" DOUBLE PRECISION NOT NULL DEFAULT 0,
                                         "created_at" TIMESTAMPTZ DEFAULT NOW(),
                                         FOREIGN KEY ("ad_id") REFERENCES "ad_campaigns" ("ad_id") ON DELETE CASCADE
);                                       -- Привязка DELETE CASCADE к ad_id

-- Добавление индексов
CREATE INDEX IF NOT EXISTS idx_history_ad_id ON "histories" ("ad_id");
CREATE INDEX IF NOT EXISTS idx_history_date ON "histories" ("date");