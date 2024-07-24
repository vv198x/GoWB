SELECT
    ac.name,
    -- p.sku,
    p.position,
    br.request
FROM
    bidder_requests br
        JOIN
    positions p ON br.id = p.request_id
        JOIN
        ad_campaigns ac ON p.sku = ac.sku
;
--WHERE
--    p.sku = 215429050;