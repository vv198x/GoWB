SELECT
    ac.ad_id,
    ac."name",
    br."request",
    ROUND((SELECT AVG(c."new_cpm")
     FROM "cpms" c
     WHERE c.ad_id = ac.ad_id
       AND c.created_at >= NOW() - INTERVAL '24 hours')) AS avg_cpm,
    ac.current_bid,
    bl."max_bid",
    p."position",
    bl."max_position",
    bl."paused"
FROM
    "bidder_lists" bl
        JOIN
    "ad_campaigns" ac ON bl."ad_id" = ac."ad_id"
        JOIN
    "bidder_requests" br ON bl."request_id" = br."id"
        JOIN
    positions p ON p.request_id = bl.request_id AND p.sku = ac.sku;