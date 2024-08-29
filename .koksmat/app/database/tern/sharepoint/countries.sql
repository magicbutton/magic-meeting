create or replace view sharepoint.countries as
select etag, 
fields->>'id' as id,
fields->>'Title' as title,
fields->>'Flag' as flag
 FROM sync.sharepoint_items('%cava3%')
where listname = 'Countries'