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
	id varchar(20) NOT NULL,
	primer_nombre text,
	otros_nombres text,
	primer_apellido text,
	segundo_apellido text,
	id_tipo_propietario smallint NOT NULL,
	CONSTRAINT propietario_pkey PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE public.propietario IS 'El registro del propietario/persona dueña del vehiculo';
-- ddl-end --
COMMENT ON COLUMN public.propietario.id IS 'Es la cedula, tarjeta de identidad seguido del numero ex CC10245032159 TI92100950407';
-- ddl-end --
COMMENT ON COLUMN public.propietario.primer_nombre IS 'Primer nombre del propietario';
-- ddl-end --
COMMENT ON COLUMN public.propietario.otros_nombres IS 'Otros nombre del propietario, ex (Maria Ana Luz, ana luz son otros nombres)';
-- ddl-end --
COMMENT ON COLUMN public.propietario.primer_apellido IS 'Primer apellido del propietario';
-- ddl-end --
COMMENT ON COLUMN public.propietario.segundo_apellido IS 'Segundo Apellido del propietario';
-- ddl-end --
COMMENT ON COLUMN public.propietario.id_tipo_propietario IS 'El propietario puede ser estudiante, funcionario, docente';
-- ddl-end --
ALTER TABLE public.propietario OWNER TO usercirce;
-- ddl-end --

-- object: public.vehiculo | type: TABLE --
-- DROP TABLE IF EXISTS public.vehiculo CASCADE;
CREATE TABLE public.vehiculo(
	id varchar(6) NOT NULL,
	id_nfc smallint NOT NULL,
	id_propietario varchar(20) NOT NULL,
	CONSTRAINT vehiculo_pkey PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE public.vehiculo IS 'El registro de la entidad vehiculo, por ahora automovil';
-- ddl-end --
COMMENT ON COLUMN public.vehiculo.id IS 'Placa del vehiculo';
-- ddl-end --
COMMENT ON COLUMN public.vehiculo.id_nfc IS 'Codigo que identifica al tag NFC';
-- ddl-end --
COMMENT ON COLUMN public.vehiculo.id_propietario IS 'El documento de identidad del propietario';
-- ddl-end --
ALTER TABLE public.vehiculo OWNER TO usercirce;
-- ddl-end --

-- object: public.tipo_propietario | type: TABLE --
-- DROP TABLE IF EXISTS public.tipo_propietario CASCADE;
CREATE TABLE public.tipo_propietario(
	id smallint NOT NULL,
	tipo varchar(25) NOT NULL,
	descripcion text,
	CONSTRAINT tipo_propietario_pkey PRIMARY KEY (id),
	CONSTRAINT tipo_propietario_nombre_uni UNIQUE (tipo)

);
-- ddl-end --
COMMENT ON TABLE public.tipo_propietario IS 'El propietario puede ser de muchos tipos que le provee de caractaristicas especiales, estos tipos inicialmente estudiante, docente, funcionario';
-- ddl-end --
ALTER TABLE public.tipo_propietario OWNER TO usercirce;
-- ddl-end --

-- object: public.isla | type: TABLE --
-- DROP TABLE IF EXISTS public.isla CASCADE;
CREATE TABLE public.isla(
	id serial NOT NULL,
	ocupado bool NOT NULL,
	id_vehiculo integer NOT NULL,
	geometria geometry(POLYGON, 4326) NOT NULL,
	hora_entrada timestamp,
	hora_salida timestamp,
	id_grupo_isla integer,
	CONSTRAINT isla_id_pkey PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE public.isla IS 'Es el espacio en que se pone el vehiculo';
-- ddl-end --
COMMENT ON COLUMN public.isla.id IS 'Autoincremental para islas, ya que por ahora no hay otra forma de identificarlos';
-- ddl-end --
COMMENT ON COLUMN public.isla.ocupado IS 'Estado que indica si el espacio esta ocupado o no';
-- ddl-end --
COMMENT ON COLUMN public.isla.id_vehiculo IS 'placa del vehiculo';
-- ddl-end --
COMMENT ON COLUMN public.isla.geometria IS 'El poligono que representa a la isla espacialmente';
-- ddl-end --
COMMENT ON COLUMN public.isla.hora_entrada IS 'Hora en que ingresa el vehiculo';
-- ddl-end --
COMMENT ON COLUMN public.isla.hora_salida IS 'Hora en que sale el vehiculo';
-- ddl-end --
COMMENT ON COLUMN public.isla.id_grupo_isla IS 'Identificador del grupo isla';
-- ddl-end --
ALTER TABLE public.isla OWNER TO usercirce;
-- ddl-end --

-- object: public.grupo_isla | type: TABLE --
-- DROP TABLE IF EXISTS public.grupo_isla CASCADE;
CREATE TABLE public.grupo_isla(
	id serial NOT NULL,
	geometria geometry(POLYGON, 4326) NOT NULL,
	id_grupo_padre integer,
	CONSTRAINT grupo_isla_id_pkey PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE public.grupo_isla IS 'Es la entidad que representa agrupacion de islas, por ejemplo podria ser Sotano 1, como un grupo de islas, o parqueadero para un grupo de sotanos';
-- ddl-end --
COMMENT ON COLUMN public.grupo_isla.id IS 'Identificador incremental ya que por el momento no hay nada mas que lo caracterice';
-- ddl-end --
COMMENT ON COLUMN public.grupo_isla.geometria IS 'Guarda el polygono que lo contiene. http://spatialreference.org/ref/epsg/4326/';
-- ddl-end --
COMMENT ON COLUMN public.grupo_isla.id_grupo_padre IS 'Se podria decir que el grupo isla tiene un padre, asi como sotano1 tiene el padre parquedero ingenieria';
-- ddl-end --
ALTER TABLE public.grupo_isla OWNER TO usercirce;
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
