-- 5_add_cpm_table.up.sql

-- Добавление нового поля subject в таблицу ad_campaigns
ALTER TABLE "ad_campaigns"
    ADD COLUMN "subject" INTEGER;

-- Создание таблицы
CREATE TABLE IF NOT EXISTS "cpms" (
                                      "id" SERIAL PRIMARY KEY,
                                      "ad_id" BIGINT NOT NULL,
                                      "old_cpm" INTEGER NOT NULL DEFAULT 0,
                                      "new_cpm" INTEGER NOT NULL DEFAULT 0,
                                      "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                      FOREIGN KEY ("ad_id") REFERENCES "ad_campaigns" ("ad_id") ON DELETE CASCADE
);

-- Создание индексов
CREATE INDEX IF NOT EXISTS idx_cpm_ad_id ON "cpms" ("ad_id");