SELECT
    ac."name",
    br."request",
    bl."current_bid",
    bl."max_bid",
    bl."paused"
FROM
    "bidder_lists" bl
        JOIN
    "ad_campaigns" ac ON bl."ad_id" = ac."ad_id"
        JOIN
    "bidder_requests" br ON bl."request_id" = br."id";