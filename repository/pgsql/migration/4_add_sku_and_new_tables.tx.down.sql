-- 4_add_sku_and_new_tables.tx.down.sql

-- Удаление нового поля sku из таблицы ad_campaigns
ALTER TABLE "ad_campaigns"
    DROP COLUMN "current_bid",
    DROP COLUMN "sku";

-- Удаление таблицы positions
DROP TABLE IF EXISTS "positions";

-- Удаление таблицы bidder_lists
DROP TABLE IF EXISTS "bidder_lists";

-- Удаление таблицы bidder_requests
DROP TABLE IF EXISTS "bidder_requests";