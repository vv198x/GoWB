select cpms.ad_id, cpms.old_cpm, cpms.new_cpm, created_at
from cpms
where cpms.created_at > now() - interval '24 hours' and cpms.ad_id = 18516064
order by created_at desc