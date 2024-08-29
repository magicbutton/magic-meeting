select listname, count(*) FROM
sync.sharepoint_items('%cava3%')
group by listname
order by listname