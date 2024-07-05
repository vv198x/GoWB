SELECT
    ac.current_bid,
    bl.max_bid,
    bl.max_position,
    bl.paused,
    p.position,
    (
        SELECT
            c.old_cpm
        FROM
            cpms c
        WHERE
            c.ad_id = bl.ad_id
        ORDER BY
            c.created_at DESC
        LIMIT 1
    ) AS old_cpm
FROM
    bidder_lists bl
        JOIN
    ad_campaigns ac ON ac.ad_id = bl.ad_id
        JOIN
    positions p ON p.sku = ac.sku
WHERE
    bl.ad_id = 17182684;