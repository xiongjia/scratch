\echo Use "CREATE EXTENSION pg_simple_ext" to load this file. \quit

CREATE FUNCTION my_simple_ext() RETURNS text
AS '$libdir/pg_simple_ext', 'simple_ext'
LANGUAGE C IMMUTABLE STRICT;

