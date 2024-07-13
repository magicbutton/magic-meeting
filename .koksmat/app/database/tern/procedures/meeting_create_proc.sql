/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- tomat sild

CREATE OR REPLACE PROCEDURE proc.create_meeting(
    p_actor_name VARCHAR,
    p_params JSONB,
    OUT p_id INTEGER
)
LANGUAGE plpgsql
AS $BODY$
DECLARE
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
         

    INSERT INTO public.meeting (
    id,
    created_at,
    updated_at,
        created_by, 
        updated_by, 
        tenant,
        searchindex,
        name,
        description,
        start,
        duration,
        location,
        status,
        exchangereference,
        exchangestatus,
        teamsreference,
        teamsstatus
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
        v_start,
        v_duration,
        v_location,
        v_status,
        v_exchangereference,
        v_exchangestatus,
        v_teamsreference,
        v_teamsstatus
    )
    RETURNING id INTO p_id;

       p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'create_meeting',
        'status', 'success',
        'description', '',
        'action', 'create_meeting',
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

  "title": "Create Meeting",
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

    -- Call the create_auditlog procedure
    CALL proc.create_auditlog(p_actor_name, p_auditlog_params, v_audit_id);
END;
$BODY$
;




