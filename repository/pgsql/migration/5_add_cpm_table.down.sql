-- 5_add_cpm_table.down.sql

ALTER TABLE "ad_campaigns"
    DROP COLUMN "subject";

-- Drop table
DROP TABLE IF EXISTS "cpms";