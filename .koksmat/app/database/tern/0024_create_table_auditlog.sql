/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.auditlog
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by character varying COLLATE pg_catalog."default"  ,

    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying COLLATE pg_catalog."default" ,

    deleted_at timestamp with time zone,
    koksmat_masterdataref VARCHAR COLLATE pg_catalog."default",
    koksmat_masterdata_id VARCHAR COLLATE pg_catalog."default",
    koksmat_masterdata_etag VARCHAR COLLATE pg_catalog."default",
    koksmat_compliancetag VARCHAR COLLATE pg_catalog."default",
    koksmat_state VARCHAR COLLATE pg_catalog."default",


    koksmat_bucket JSONB 

    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,searchindex character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,action character varying COLLATE pg_catalog."default"  NOT NULL
    ,user_id int   NOT NULL
    ,entity character varying COLLATE pg_catalog."default"  NOT NULL
    ,entityid character varying COLLATE pg_catalog."default"  NOT NULL
    ,timestamp character varying COLLATE pg_catalog."default"   NOT NULL


);

                ALTER TABLE IF EXISTS public.auditlog
                ADD FOREIGN KEY (user_id)
                REFERENCES public.user (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----

DROP TABLE public.auditlog;

