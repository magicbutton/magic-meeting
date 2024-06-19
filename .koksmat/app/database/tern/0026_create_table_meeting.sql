/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   


-- sure sild

CREATE TABLE public.meeting
(
    id SERIAL PRIMARY KEY,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by character varying COLLATE pg_catalog."default"  ,

    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by character varying COLLATE pg_catalog."default" ,

    deleted_at timestamp with time zone
    ,tenant character varying COLLATE pg_catalog."default"  NOT NULL
    ,searchindex character varying COLLATE pg_catalog."default"  NOT NULL
    ,name character varying COLLATE pg_catalog."default"  NOT NULL
    ,description character varying COLLATE pg_catalog."default" 
    ,start character varying COLLATE pg_catalog."default"   NOT NULL
    ,duration character varying COLLATE pg_catalog."default"   NOT NULL
    ,location character varying COLLATE pg_catalog."default"  NOT NULL
    ,status character varying COLLATE pg_catalog."default"  NOT NULL
    ,exchangereference character varying COLLATE pg_catalog."default" 
    ,exchangestatus character varying COLLATE pg_catalog."default" 
    ,teamsreference character varying COLLATE pg_catalog."default" 
    ,teamsstatus character varying COLLATE pg_catalog."default" 


);

                -- lollipop
                CREATE TABLE public.meeting_m2m_serviceorder (
                id SERIAL PRIMARY KEY,
                created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                created_by character varying COLLATE pg_catalog."default"  ,
                updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
                updated_by character varying COLLATE pg_catalog."default",
                deleted_at timestamp with time zone
                
                    ,meeting_id int  
 
                    ,serviceorder_id int  
 

                );
            

                ALTER TABLE public.meeting_m2m_serviceorder
                ADD FOREIGN KEY (meeting_id)
                REFERENCES public.meeting (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;

                ALTER TABLE public.meeting_m2m_serviceorder
                ADD FOREIGN KEY (serviceorder_id)
                REFERENCES public.serviceorder (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
                NOT VALID;


---- create above / drop below ----
DROP TABLE IF EXISTS public.meeting_m2m_serviceorder;
DROP TABLE public.meeting;

