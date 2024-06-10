-- 2_add_balance_table.tx.down.sql

-- Удаление таблицы balance
DROP TABLE IF EXISTS "balances";

-- Удаление поля do_not_refill из таблицы ad_campaigns
ALTER TABLE "ad_campaigns"
    DROP COLUMN "do_not_refill";