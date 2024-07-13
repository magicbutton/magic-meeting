/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: true
---
*/   


-- tomat sild

CREATE OR REPLACE PROCEDURE proc.create_auditlog(
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
    v_action VARCHAR;
    v_user_id INTEGER;
    v_entity VARCHAR;
    v_entityid VARCHAR;
    v_timestamp TIMESTAMP WITH TIME ZONE;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;

BEGIN
    v_tenant := p_params->>'tenant';
    v_searchindex := p_params->>'searchindex';
    v_name := p_params->>'name';
    v_description := p_params->>'description';
    v_action := p_params->>'action';
    v_user_id := p_params->>'user_id';
    v_entity := p_params->>'entity';
    v_entityid := p_params->>'entityid';
    v_timestamp := p_params->>'timestamp';
         

    INSERT INTO public.auditlog (
    id,
    created_at,
    updated_at,
        created_by, 
        updated_by, 
        tenant,
        searchindex,
        name,
        description,
        action,
        user_id,
        entity,
        entityid,
        timestamp
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
        v_action,
        v_user_id,
        v_entity,
        v_entityid,
        v_timestamp
    )
    RETURNING id INTO p_id;

       p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'create_auditlog',
        'status', 'success',
        'description', '',
        'action', 'create_auditlog',
        'entity', 'auditlog',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
{
   "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://booking.services.koksmat.com/.schema.json",
   
  "type": "object",

  "title": "Create AuditLog",
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
    "action": { 
    "type": "string",
    "description":"" },
    "user_id": { 
    "type": "number",
    "description":"" },
    "entity": { 
    "type": "string",
    "description":"" },
    "entityid": { 
    "type": "string",
    "description":"" },
    "timestamp": { 
    "type": "string",
    "description":"" }

    }
}

##MAGICAPP-END##*/

    -- Call the create_auditlog procedure
    -- CALL proc.create_auditlog(p_actor_name, p_auditlog_params, v_audit_id);
END;
$BODY$
;




