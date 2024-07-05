SELECT
    br.request,
    p.sku,
    p.position
FROM
    bidder_requests br
        JOIN
    positions p ON br.id = p.request_id;
--WHERE
--    p.sku = 215429050;