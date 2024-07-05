INSERT INTO "bidder_lists" ("ad_id", "request_id", "max_bid", "paused", "created_at", "updated_at")
SELECT "ad_id",
       1 AS "request_id",
       0 AS "max_bid",
       TRUE AS "paused",
       now() AS "created_at",
       now() AS "updated_at"
FROM "ad_campaigns";

-- Переписать на обновление