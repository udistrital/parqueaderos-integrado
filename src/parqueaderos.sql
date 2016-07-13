-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.8.1
-- PostgreSQL version: 9.4
-- Project Site: pgmodeler.com.br
-- Model Author: ---


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: new_database | type: DATABASE --
-- -- DROP DATABASE IF EXISTS new_database;
-- CREATE DATABASE new_database
-- ;
-- -- ddl-end --
-- 

-- object: public.propietario | type: TABLE --
-- DROP TABLE IF EXISTS public.propietario CASCADE;
CREATE TABLE public.propietario(
	id serial NOT NULL,
	documento text NOT NULL,
	primer_nombre text,
	segundo_nombre text,
	primer_apellido text,
	segundo_apellido text,
	id_tipo_propietario integer NOT NULL,
	CONSTRAINT propietario_pkey PRIMARY KEY (id),
	CONSTRAINT propietario_documento_uni UNIQUE (documento)

);
-- ddl-end --
COMMENT ON COLUMN public.propietario.documento IS 'CC1026281485';
-- ddl-end --
ALTER TABLE public.propietario OWNER TO postgres;
-- ddl-end --

-- object: public.vehiculo | type: TABLE --
-- DROP TABLE IF EXISTS public.vehiculo CASCADE;
CREATE TABLE public.vehiculo(
	id serial NOT NULL,
	id_nfc smallint NOT NULL,
	placa text NOT NULL,
	id_propietario integer NOT NULL,
	CONSTRAINT vehiculo_pkey PRIMARY KEY (id),
	CONSTRAINT vehiculo_placa_uni UNIQUE (placa)

);
-- ddl-end --
ALTER TABLE public.vehiculo OWNER TO postgres;
-- ddl-end --

-- object: public.tipo_propietario | type: TABLE --
-- DROP TABLE IF EXISTS public.tipo_propietario CASCADE;
CREATE TABLE public.tipo_propietario(
	id serial NOT NULL,
	nombre text NOT NULL,
	descripcion text,
	CONSTRAINT tipo_propietario_pkey PRIMARY KEY (id),
	CONSTRAINT tipo_propietario_nombre_uni UNIQUE (nombre)

);
-- ddl-end --
ALTER TABLE public.tipo_propietario OWNER TO postgres;
-- ddl-end --

-- object: public.isla | type: TABLE --
-- DROP TABLE IF EXISTS public.isla CASCADE;
CREATE TABLE public.isla(
	id serial NOT NULL,
	estado bool NOT NULL,
	id_vehiculo integer NOT NULL,
	geometria polygon NOT NULL,
	hora_entrada timestamp,
	hora_salida timestamp,
	id_grupo_isla integer,
	CONSTRAINT isla_id_pkey PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.isla OWNER TO postgres;
-- ddl-end --

-- object: public.grupo_isla | type: TABLE --
-- DROP TABLE IF EXISTS public.grupo_isla CASCADE;
CREATE TABLE public.grupo_isla(
	id serial NOT NULL,
	geometria polygon NOT NULL,
	id_grupo_padre integer,
	CONSTRAINT grupo_isla_id_pkey PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE public.grupo_isla OWNER TO postgres;
-- ddl-end --

-- object: propietario_tipo_propietario_id_fkey | type: CONSTRAINT --
-- ALTER TABLE public.propietario DROP CONSTRAINT IF EXISTS propietario_tipo_propietario_id_fkey CASCADE;
ALTER TABLE public.propietario ADD CONSTRAINT propietario_tipo_propietario_id_fkey FOREIGN KEY (id_tipo_propietario)
REFERENCES public.tipo_propietario (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: vehiculo_propietario_id_propietario_fkey | type: CONSTRAINT --
-- ALTER TABLE public.vehiculo DROP CONSTRAINT IF EXISTS vehiculo_propietario_id_propietario_fkey CASCADE;
ALTER TABLE public.vehiculo ADD CONSTRAINT vehiculo_propietario_id_propietario_fkey FOREIGN KEY (id_propietario)
REFERENCES public.propietario (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: isla_vehiculo_id_vehiculo_fkey | type: CONSTRAINT --
-- ALTER TABLE public.isla DROP CONSTRAINT IF EXISTS isla_vehiculo_id_vehiculo_fkey CASCADE;
ALTER TABLE public.isla ADD CONSTRAINT isla_vehiculo_id_vehiculo_fkey FOREIGN KEY (id_vehiculo)
REFERENCES public.vehiculo (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: isla_grupo_isla_id_grupo_isla_fkey | type: CONSTRAINT --
-- ALTER TABLE public.isla DROP CONSTRAINT IF EXISTS isla_grupo_isla_id_grupo_isla_fkey CASCADE;
ALTER TABLE public.isla ADD CONSTRAINT isla_grupo_isla_id_grupo_isla_fkey FOREIGN KEY (id_grupo_isla)
REFERENCES public.grupo_isla (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


