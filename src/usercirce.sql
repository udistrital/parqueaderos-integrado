CREATE ROLE usercirce WITH 
	INHERIT
	LOGIN
	ENCRYPTED PASSWORD 'oascirce';

DROP DATABASE IF EXISTS circe;
CREATE DATABASE circe;
\connect circe;
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_topology;
CREATE EXTENSION ogr_fdw;
SELECT postgis_full_version();
ALTER DATABASE circe OWNER TO usercirce;
