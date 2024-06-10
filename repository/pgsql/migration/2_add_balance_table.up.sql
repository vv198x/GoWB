-- 2_add_balance_table.tx.up.sql

-- Создание таблицы balance
CREATE TABLE IF NOT EXISTS "balances" (
                                          "ad_id" BIGINT PRIMARY KEY REFERENCES "ad_campaigns" ("ad_id") ON DELETE CASCADE,
                                          "balance" DOUBLE PRECISION,
                                          "updated_at" TIMESTAMPTZ DEFAULT now()
);
-- Добавление нового поля do_not_refill в таблицу ad_campaigns
ALTER TABLE "ad_campaigns"
    ADD COLUMN "do_not_refill" BOOLEAN DEFAULT FALSE;