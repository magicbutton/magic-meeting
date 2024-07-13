/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- krydder sild

CREATE OR REPLACE PROCEDURE proc.delete_productionorder(
    p_actor_name VARCHAR,
    p_params JSONB
)
LANGUAGE plpgsql
AS $BODY$
DECLARE
    v_id INTEGER;
    v_hard BOOLEAN;
     v_rows_updated INTEGER;
        v_audit_id integer;  -- Variable to hold the OUT parameter value
    p_auditlog_params jsonb;


BEGIN
    v_id := p_params->>'id';
    v_hard := p_params->>'hard';
  
    IF v_hard THEN
     DELETE FROM public.productionorder
        WHERE id = v_id;
       
    ELSE
        UPDATE public.productionorder
        SET deleted_at = CURRENT_TIMESTAMP,
            updated_at = CURRENT_TIMESTAMP,
            updated_by = p_actor_name
        WHERE id = v_id;
        
        GET DIAGNOSTICS v_rows_updated = ROW_COUNT;
    
        IF v_rows_updated < 1 THEN
            RAISE EXCEPTION 'No records updated. productionorder ID % not found', v_id ;
        END IF;
    END IF;
           p_auditlog_params := jsonb_build_object(
        'tenant', '',
        'searchindex', '',
        'name', 'delete_productionorder',
        'status', 'success',
        'description', '',
        'action', 'delete_productionorder',
        'entity', 'productionorder',
        'entityid', -1,
        'actor', p_actor_name,
        'metadata', p_params
    );
/*###MAGICAPP-START##
{
     "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://booking.services.koksmat.com/.schema.json",
   
  "type": "object",

  "title": "Delete ProductionOrder",
  "description": "Delete operation",
  "properties": {
   "id": { "type": "number" },
    "hard": { "type": "boolean" }

    }
}
##MAGICAPP-END##*/
END;
$BODY$
;
