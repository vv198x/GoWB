SELECT
    ac.ad_id,
    ac."name",
    br."request",
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