-- drop view sharepoint.floors
create or replace view sharepoint.rooms as
select etag, 
fields->>'id' as id,
fields->>'Title' as Title,
fields->>'Email' as Email,
fields->>'Capacity' as Capacity,
fields->>'Provisioning_x0020_Status' as Provisioning_x0020_Status,
fields->>'FloorLookupId' as FloorLookupId
-- *
-- fields->>'State' as State,
-- fields->>'Street' as Street,
-- fields->>'Latitude' as Latitude,
-- fields->>'Longitude' as Longitude,
-- fields->>'PostalCode' as PostalCode
 FROM sync.sharepoint_items('%cava3%')
where listname = 'Rooms'