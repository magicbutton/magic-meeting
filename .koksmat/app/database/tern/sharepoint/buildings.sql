create or replace view sharepoint.sites as
select etag, 
fields->>'id' as id,
fields->>'Title' as Title,
fields->>'State' as State,
fields->>'Street' as Street,
fields->>'Latitude' as Latitude,
fields->>'Longitude' as Longitude,
fields->>'PostalCode' as PostalCode
 FROM sync.sharepoint_items('%cava3%')
where listname = 'Locations'