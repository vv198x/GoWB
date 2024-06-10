SELECT
    ac.name,
    COALESCE(b.balance, 0) AS balance,
    CASE
        WHEN ac.status = 9 THEN 'RUN'
        WHEN ac.status = 11 THEN 'PAUSE'
        ELSE 'UNKNOWN'
        END AS status,
    ac.budget,
    NOT ac.do_not_refill AS refill
FROM
    ad_campaigns ac
        LEFT JOIN
    balances b ON ac.ad_id = b.ad_id;