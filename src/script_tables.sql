-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.8.1
-- PostgreSQL version: 9.4
-- Project Site: pgmodeler.com.br
-- Model Author: ---

-- object: usercirce | type: ROLE --
-- DROP ROLE IF EXISTS usercirce;
CREATE ROLE usercirce WITH 
	INHERIT
	LOGIN
	ENCRYPTED PASSWORD '********';
-- ddl-end --


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: udistrital | type: DATABASE --
-- -- DROP DATABASE IF EXISTS udistrital;
-- CREATE DATABASE udistrital
-- 	ENCODING = 'UTF8'
-- 	LC_COLLATE = 'es_CO.UTF8'
-- 	LC_CTYPE = 'es_CO.UTF8'
-- 	TABLESPACE = pg_default
-- 	OWNER = usercirce
-- ;
-- -- ddl-end --
-- 

-- object: parqueaderos | type: SCHEMA --
-- DROP SCHEMA IF EXISTS parqueaderos CASCADE;
CREATE SCHEMA parqueaderos;
-- ddl-end --
ALTER SCHEMA parqueaderos OWNER TO postgres;
-- ddl-end --

-- object: topology | type: SCHEMA --
-- DROP SCHEMA IF EXISTS topology CASCADE;
CREATE SCHEMA topology;
-- ddl-end --
ALTER SCHEMA topology OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,parqueaderos,topology;
-- ddl-end --

-- object: postgis | type: EXTENSION --
-- DROP EXTENSION IF EXISTS postgis CASCADE;
CREATE EXTENSION postgis
      WITH SCHEMA parqueaderos
      VERSION '2.2.1';
-- ddl-end --
COMMENT ON EXTENSION postgis IS 'PostGIS geometry, geography, and raster spatial types and functions';
-- ddl-end --

-- object: postgis_topology | type: EXTENSION --
-- DROP EXTENSION IF EXISTS postgis_topology CASCADE;
CREATE EXTENSION postgis_topology
      WITH SCHEMA topology
      VERSION '2.2.1';
-- ddl-end --
COMMENT ON EXTENSION postgis_topology IS 'PostGIS topology spatial types and functions';
-- ddl-end --

-- object: parqueaderos.propietario_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS parqueaderos.propietario_id_seq CASCADE;
CREATE SEQUENCE parqueaderos.propietario_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE parqueaderos.propietario_id_seq OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.propietario | type: TABLE --
-- DROP TABLE IF EXISTS parqueaderos.propietario CASCADE;
CREATE TABLE parqueaderos.propietario(
	id integer NOT NULL DEFAULT nextval('parqueaderos.propietario_id_seq'::regclass),
	documento character varying(20) NOT NULL,
	primer_nombre text,
	otros_nombres text,
	primer_apellido text,
	segundo_apellido text,
	id_tipo_propietario integer NOT NULL,
	CONSTRAINT propietario_pkey PRIMARY KEY (id),
	CONSTRAINT propietario_documento_uni UNIQUE (documento)

);
-- ddl-end --
COMMENT ON TABLE parqueaderos.propietario IS 'El registro del propietario/persona dueña del vehiculo';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.propietario.documento IS 'Es la cedula, tarjeta de identidad seguido del numero ex CC10245032159 TI92100950407';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.propietario.primer_nombre IS 'Primer nombre del propietario';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.propietario.otros_nombres IS 'Otros nombre del propietario, ex (Maria Ana Luz, ana luz son otros nombres)';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.propietario.primer_apellido IS 'Primer apellido del propietario';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.propietario.segundo_apellido IS 'Segundo Apellido del propietario';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.propietario.id_tipo_propietario IS 'El propietario puede ser estudiante, funcionario, docente';
-- ddl-end --
ALTER TABLE parqueaderos.propietario OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.vehiculo_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS parqueaderos.vehiculo_id_seq CASCADE;
CREATE SEQUENCE parqueaderos.vehiculo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE parqueaderos.vehiculo_id_seq OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.vehiculo | type: TABLE --
-- DROP TABLE IF EXISTS parqueaderos.vehiculo CASCADE;
CREATE TABLE parqueaderos.vehiculo(
	id integer NOT NULL DEFAULT nextval('parqueaderos.vehiculo_id_seq'::regclass),
	placa character varying(6) NOT NULL,
	id_nfc bigint NOT NULL,
	id_propietario integer NOT NULL,
	CONSTRAINT vehiculo_pkey PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE parqueaderos.vehiculo IS 'El registro de la entidad vehiculo, por ahora automovil';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.vehiculo.placa IS 'Placa del vehiculo';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.vehiculo.id_nfc IS 'Codigo que identifica al tag NFC';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.vehiculo.id_propietario IS 'El documento de identidad del propietario';
-- ddl-end --
ALTER TABLE parqueaderos.vehiculo OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.tipo_propietario_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS parqueaderos.tipo_propietario_id_seq CASCADE;
CREATE SEQUENCE parqueaderos.tipo_propietario_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE parqueaderos.tipo_propietario_id_seq OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.tipo_propietario | type: TABLE --
-- DROP TABLE IF EXISTS parqueaderos.tipo_propietario CASCADE;
CREATE TABLE parqueaderos.tipo_propietario(
	id integer NOT NULL DEFAULT nextval('parqueaderos.tipo_propietario_id_seq'::regclass),
	tipo character varying(25) NOT NULL,
	descripcion text,
	CONSTRAINT tipo_propietario_pkey PRIMARY KEY (id),
	CONSTRAINT tipo_propietario_tipo_uni UNIQUE (tipo)

);
-- ddl-end --
COMMENT ON TABLE parqueaderos.tipo_propietario IS 'El propietario puede ser de muchos tipos que le provee de caractaristicas especiales, estos tipos inicialmente estudiante, docente, funcionario';
-- ddl-end --
ALTER TABLE parqueaderos.tipo_propietario OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.grupo_isla_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS parqueaderos.grupo_isla_id_seq CASCADE;
CREATE SEQUENCE parqueaderos.grupo_isla_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE parqueaderos.grupo_isla_id_seq OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.grupo_isla | type: TABLE --
-- DROP TABLE IF EXISTS parqueaderos.grupo_isla CASCADE;
CREATE TABLE parqueaderos.grupo_isla(
	id integer NOT NULL DEFAULT nextval('parqueaderos.grupo_isla_id_seq'::regclass),
	geometria geometry NOT NULL,
	id_grupo_padre integer,
	CONSTRAINT grupo_isla_id_pkey PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE parqueaderos.grupo_isla IS 'Es la entidad que representa agrupacion de islas, por ejemplo podria ser Sotano 1, como un grupo de islas, o parqueadero para un grupo de sotanos';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.grupo_isla.id IS 'Identificador incremental ya que por el momento no hay nada mas que lo caracterice';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.grupo_isla.geometria IS 'Guarda el polygono que lo contiene. http://spatialreference.org/ref/epsg/4326/';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.grupo_isla.id_grupo_padre IS 'Se podria decir que el grupo isla tiene un padre, asi como sotano1 tiene el padre parquedero ingenieria';
-- ddl-end --
ALTER TABLE parqueaderos.grupo_isla OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.migrations_status | type: TYPE --
-- DROP TYPE IF EXISTS parqueaderos.migrations_status CASCADE;
CREATE TYPE parqueaderos.migrations_status AS
 ENUM ('update','rollback');
-- ddl-end --
ALTER TYPE parqueaderos.migrations_status OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.migrations_id_migration_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS parqueaderos.migrations_id_migration_seq CASCADE;
CREATE SEQUENCE parqueaderos.migrations_id_migration_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE parqueaderos.migrations_id_migration_seq OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.migrations | type: TABLE --
-- DROP TABLE IF EXISTS parqueaderos.migrations CASCADE;
CREATE TABLE parqueaderos.migrations(
	id_migration integer NOT NULL DEFAULT nextval('parqueaderos.migrations_id_migration_seq'::regclass),
	name character varying(255) DEFAULT NULL::character varying,
	created_at timestamp NOT NULL DEFAULT now(),
	statements text,
	rollback_statements text,
	status parqueaderos.migrations_status,
	CONSTRAINT migrations_pkey PRIMARY KEY (id_migration)

);
-- ddl-end --
ALTER TABLE parqueaderos.migrations OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.registro_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS parqueaderos.registro_id_seq CASCADE;
CREATE SEQUENCE parqueaderos.registro_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE parqueaderos.registro_id_seq OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.registro | type: TABLE --
-- DROP TABLE IF EXISTS parqueaderos.registro CASCADE;
CREATE TABLE parqueaderos.registro(
	id integer NOT NULL DEFAULT nextval('parqueaderos.registro_id_seq'::regclass),
	id_vehiculo integer NOT NULL,
	id_isla integer NOT NULL,
	hora_entrada timestamp,
	hora_salida timestamp,
	CONSTRAINT registro_pkey PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE parqueaderos.registro IS 'Es el espacio donde queda almacenado el registro de entrada y salida de un vehiculo al parqueadero';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.registro.id IS 'Autoincremental para registros, ya que por ahora no hay otra forma de identificarlos';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.registro.id_vehiculo IS 'identificador del vehiculo y del propietario';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.registro.id_isla IS 'identificador de la isla';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.registro.hora_entrada IS 'Hora en que ingresa el vehiculo';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.registro.hora_salida IS 'Hora en que sale el vehiculo';
-- ddl-end --
ALTER TABLE parqueaderos.registro OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.isla_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS parqueaderos.isla_id_seq CASCADE;
CREATE SEQUENCE parqueaderos.isla_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
ALTER SEQUENCE parqueaderos.isla_id_seq OWNER TO usercirce;
-- ddl-end --

-- object: parqueaderos.isla | type: TABLE --
-- DROP TABLE IF EXISTS parqueaderos.isla CASCADE;
CREATE TABLE parqueaderos.isla(
	id integer NOT NULL DEFAULT nextval('parqueaderos.isla_id_seq'::regclass),
	ocupado boolean NOT NULL,
	geometria geometry NOT NULL,
	id_grupo_isla integer,
	CONSTRAINT isla_id_pkey PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE parqueaderos.isla IS 'Es el espacio en que se pone la isla';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.isla.id IS 'Autoincremental para islas, ya que por ahora no hay otra forma de identificarlos';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.isla.ocupado IS 'Estado que indica si el espacio esta ocupado o no';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.isla.geometria IS 'El poligono que representa a la isla espacialmente';
-- ddl-end --
COMMENT ON COLUMN parqueaderos.isla.id_grupo_isla IS 'Identificador del grupo isla';
-- ddl-end --
ALTER TABLE parqueaderos.isla OWNER TO usercirce;
-- ddl-end --

-- object: propietario_tipo_propietario_id_fkey | type: CONSTRAINT --
-- ALTER TABLE parqueaderos.propietario DROP CONSTRAINT IF EXISTS propietario_tipo_propietario_id_fkey CASCADE;
ALTER TABLE parqueaderos.propietario ADD CONSTRAINT propietario_tipo_propietario_id_fkey FOREIGN KEY (id_tipo_propietario)
REFERENCES parqueaderos.tipo_propietario (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: vehiculo_propietario_id_propietario_fkey | type: CONSTRAINT --
-- ALTER TABLE parqueaderos.vehiculo DROP CONSTRAINT IF EXISTS vehiculo_propietario_id_propietario_fkey CASCADE;
ALTER TABLE parqueaderos.vehiculo ADD CONSTRAINT vehiculo_propietario_id_propietario_fkey FOREIGN KEY (id_propietario)
REFERENCES parqueaderos.propietario (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: registro_vehiculo_id_vehiculo_fkey | type: CONSTRAINT --
-- ALTER TABLE parqueaderos.registro DROP CONSTRAINT IF EXISTS registro_vehiculo_id_vehiculo_fkey CASCADE;
ALTER TABLE parqueaderos.registro ADD CONSTRAINT registro_vehiculo_id_vehiculo_fkey FOREIGN KEY (id_vehiculo)
REFERENCES parqueaderos.vehiculo (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: registro_isla_id_isla_fkey | type: CONSTRAINT --
-- ALTER TABLE parqueaderos.registro DROP CONSTRAINT IF EXISTS registro_isla_id_isla_fkey CASCADE;
ALTER TABLE parqueaderos.registro ADD CONSTRAINT registro_isla_id_isla_fkey FOREIGN KEY (id_isla)
REFERENCES parqueaderos.isla (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: isla_grupo_isla_id_grupo_isla_fkey | type: CONSTRAINT --
-- ALTER TABLE parqueaderos.isla DROP CONSTRAINT IF EXISTS isla_grupo_isla_id_grupo_isla_fkey CASCADE;
ALTER TABLE parqueaderos.isla ADD CONSTRAINT isla_grupo_isla_id_grupo_isla_fkey FOREIGN KEY (id_grupo_isla)
REFERENCES parqueaderos.grupo_isla (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


