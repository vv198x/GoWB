SELECT
    ac.ad_id,
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
WHERE
    ac.type = 9
;
--
--    p.sku = 215429050;