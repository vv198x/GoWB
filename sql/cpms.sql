select cpms.ad_id, cpms.old_cpm, cpms.new_cpm, cpms.old_position, cpms.created_at
from cpms
where cpms.created_at > now() - interval '24 hours' and cpms.ad_id = 17182914
order by created_at desc

-- 17182684 лкар ап 150 а там
-- 19647147  вишня коллаген 150
-- 18450096 алкар буст апельсин
-- 17182977 жги
-- 17741485 жги ананас
-- 18516064 фл5
-- 17182914 ФЛ-0009