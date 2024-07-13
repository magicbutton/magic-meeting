/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sherry sild

CREATE OR REPLACE PROCEDURE proc.update_serviceorder(
    p_actor_name VARCHAR,
    p_params JSONB
)
LANGUAGE plpgsql
AS $BODY$
DECLARE
    v_id INTEGER;
       v_rows_updated INTEGER;
v_tenant VARCHAR COLLATE pg_catalog."default" ;
    v_searchindex VARCHAR COLLATE pg_catalog."default" ;
    v_name VARCHAR COLLATE pg_catalog."default" ;
    v_description VARCHAR COLLATE pg_catalog."default";
    v_deliverydate TIMESTAMP WITH TIME ZONE;
    v_deliverto_id INTEGER;
    v_status VARCHAR;
    v_payment_id INTEGER;
    v_orderdata JSONB;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

    
BEGIN
    v_id := p_params->>'id';
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_deliverydate := p_params->>'deliverydate';
    v_deliverto_id := p_params->>'deliverto_id';
    v_status := p_params->>'status';
    v_payment_id := p_params->>'payment_id';
    v_orderdata := p_params->>'orderdata';
         
    
        
    UPDATE public.serviceorder
    SET updated_by = p_actor_name,
        updated_at = CURRENT_TIMESTAMP,
        tenant = v_tenant,
        searchindex = v_searchindex,
        name = v_name,
        description = v_description,
        deliverydate = v_deliverydate,
        deliverto_id = v_deliverto_id,
        status = v_status,
        payment_id = v_payment_id,
        orderdata = v_orderdata
    WHERE id = v_id;

    GET DIAGNOSTICS v_rows_updated = ROW_COUNT;
    
    IF v_rows_updated < 1 THEN
        RAISE EXCEPTION 'No records updated. serviceorder ID % not found', v_id ;
    END IF;

           p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'update_serviceorder',
        'status', 'success',
        'description', '',
        'action', 'update_serviceorder',
        'entity', 'serviceorder',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
{
   "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://booking.services.koksmat.com/.schema.json",
   
  "type": "object",

  "properties": {
    "title": "Update ServiceOrder",
  "description": "Update operation",
  
    "tenant": { 
    "type": "string",
    "description":"" },
    "searchindex": { 
    "type": "string",
    "description":"Search Index is used for concatenating all searchable fields in a single field making in easier to search\n" },
    "name": { 
    "type": "string",
    "description":"" },
    "description": { 
    "type": "string",
    "description":"" },
    "deliverydate": { 
    "type": "string",
    "description":"" },
    "deliverto_id": { 
    "type": "number",
    "description":"" },
    "status": { 
    "type": "string",
    "description":"" },
    "payment_id": { 
    "type": "number",
    "description":"" },
    "orderdata": { 
    "type": "object",
    "description":"" }

    }
}
##MAGICAPP-END##*/
END;
$BODY$
;


