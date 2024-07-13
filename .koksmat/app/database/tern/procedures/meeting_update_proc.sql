/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sherry sild

CREATE OR REPLACE PROCEDURE proc.update_meeting(
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
    v_start TIMESTAMP WITH TIME ZONE;
    v_duration INTEGER;
    v_location VARCHAR;
    v_status VARCHAR;
    v_exchangereference VARCHAR;
    v_exchangestatus VARCHAR;
    v_teamsreference VARCHAR;
    v_teamsstatus VARCHAR;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

    
BEGIN
    v_id := p_params->>'id';
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_start := p_params->>'start';
    v_duration := p_params->>'duration';
    v_location := p_params->>'location';
    v_status := p_params->>'status';
    v_exchangereference := p_params->>'exchangereference';
    v_exchangestatus := p_params->>'exchangestatus';
    v_teamsreference := p_params->>'teamsreference';
    v_teamsstatus := p_params->>'teamsstatus';
         
    
        
    UPDATE public.meeting
    SET updated_by = p_actor_name,
        updated_at = CURRENT_TIMESTAMP,
        tenant = v_tenant,
        searchindex = v_searchindex,
        name = v_name,
        description = v_description,
        start = v_start,
        duration = v_duration,
        location = v_location,
        status = v_status,
        exchangereference = v_exchangereference,
        exchangestatus = v_exchangestatus,
        teamsreference = v_teamsreference,
        teamsstatus = v_teamsstatus
    WHERE id = v_id;

    GET DIAGNOSTICS v_rows_updated = ROW_COUNT;
    
    IF v_rows_updated < 1 THEN
        RAISE EXCEPTION 'No records updated. meeting ID % not found', v_id ;
    END IF;

           p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'update_meeting',
        'status', 'success',
        'description', '',
        'action', 'update_meeting',
        'entity', 'meeting',
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
    "title": "Update Meeting",
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
    "start": { 
    "type": "string",
    "description":"" },
    "duration": { 
    "type": "number",
    "description":"Duration in minutes" },
    "location": { 
    "type": "string",
    "description":"" },
    "status": { 
    "type": "string",
    "description":"" },
    "exchangereference": { 
    "type": "string",
    "description":"" },
    "exchangestatus": { 
    "type": "string",
    "description":"" },
    "teamsreference": { 
    "type": "string",
    "description":"" },
    "teamsstatus": { 
    "type": "string",
    "description":"" }

    }
}
##MAGICAPP-END##*/
END;
$BODY$
;


