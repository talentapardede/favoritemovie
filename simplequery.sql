select p.id, p.user_name, p2.user_name as parent_name 
from public.parents p 
left join public.parents p2 on p.parent = p2.id;