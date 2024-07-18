SELECT
    ac.name,
    h.date,
    SUM(h.amount) AS total_amount
FROM
    histories h
        JOIN
    ad_campaigns ac ON h.ad_id = ac.ad_id
WHERE
    h.date >= CURRENT_DATE - INTERVAL '2 days'
GROUP BY
    ac.name,
    h.date
ORDER BY
    h.date DESC,
    ac.name;
