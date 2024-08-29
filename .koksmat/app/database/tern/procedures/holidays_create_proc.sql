/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- tomat sild

CREATE OR REPLACE FUNCTION proc.create_holidays(
    p_actor_name VARCHAR,
    p_params JSONB
   
)
RETURNS JSONB LANGUAGE plpgsql 
AS $$
DECLARE
       v_rows_updated INTEGER;
v_tenant VARCHAR COLLATE pg_catalog."default" ;
    v_searchindex VARCHAR COLLATE pg_catalog."default" ;
    v_name VARCHAR COLLATE pg_catalog."default" ;
    v_description VARCHAR COLLATE pg_catalog."default";
    v_country_id INTEGER;
    v_fromhour VARCHAR;
    v_tohour VARCHAR;
    v_date TIMESTAMP WITH TIME ZONE;
    v_id INTEGER;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

BEGIN
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_country_id := p_params->>'country_id';
    v_fromhour := p_params->>'fromhour';
    v_tohour := p_params->>'tohour';
    v_date := p_params->>'date';
         

    INSERT INTO public.holidays (
    id,
    created_at,
    updated_at,
        created_by, 
        updated_by, 
        tenant,
        searchindex,
        name,
        description,
        country_id,
        fromhour,
        tohour,
        date
    )
    VALUES (
        DEFAULT,
        DEFAULT,
        DEFAULT,
        p_actor_name, 
        p_actor_name,  -- Use the same value for updated_by
        v_tenant,
        v_searchindex,
        v_name,
        v_description,
        v_country_id,
        v_fromhour,
        v_tohour,
        v_date
    )
    RETURNING id INTO v_id;

       p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'create_holidays',
        'status', 'success',
        'description', '',
        'action', 'create_holidays',
        'entity', 'holidays',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
{
   "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://booking.services.koksmat.com/.schema.json",
   
  "type": "object",

  "title": "Create Holidays",
  "description": "Create operation",

  "properties": {
  
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
    "country_id": { 
    "type": "number",
    "description":"" },
    "fromhour": { 
    "type": "string",
    "description":"" },
    "tohour": { 
    "type": "string",
    "description":"" },
    "date": { 
    "type": "string",
    "description":"" }

    }
}

##MAGICAPP-END##*/

    -- Call the create_auditlog procedure
    CALL proc.create_auditlog(p_actor_name, p_auditlog_params, v_audit_id);

    return jsonb_build_object(
    'comment','created',
    'id',v_id);

END;
$$ 
;




