create schema if not exists sharepoint;drop view if exists sharepoint.cava3 cascade;create or replace view sharepoint.cava3 as
select 
*
from (
select 
--*,

batchname,
detail_element -> 'data' ->> 'listname' as listname, 
detail_element -> 'data' ->> 'field' as field


FROM (
    SELECT
        importdata.name AS batchname,
        jsonb_array_elements(data) AS detail_element    FROM
        importdata
    WHERE
        importdata.name::text = 'mix|sharepoint.cava3'::text
) AS subquery
) as xx


--LIMIT 100;
