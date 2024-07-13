/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
 */
-- sure sild
CREATE TABLE
    cache.microsoftgraph (
        id SERIAL PRIMARY KEY,
        created_at timestamp
        with
            time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
            created_by character varying COLLATE pg_catalog."default",
            name character varying COLLATE pg_catalog."default" NOT NULL,
            url character varying COLLATE pg_catalog."default" NOT NULL,
            upn character varying COLLATE pg_catalog."default" NOT NULL,
            data JSONB NOT NULL
    );