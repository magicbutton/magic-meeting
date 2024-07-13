drop schema if exists "cache";
-- cascade;

CREATE SCHEMA "cache";

{{ template "microsoftgraph.sql".}} 