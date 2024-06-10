SELECT
    ac.ad_id
FROM
    ad_campaigns ac
        LEFT JOIN
    balances b ON ac.ad_id = b.ad_id
WHERE
        COALESCE(b.balance, 0) < 500
  AND ac.do_not_refill = FALSE;