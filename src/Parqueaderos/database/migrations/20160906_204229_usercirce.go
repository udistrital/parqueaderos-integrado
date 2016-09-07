package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Usercirce_20160906_204229 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Usercirce_20160906_204229{}
	m.Created = "20160906_204229"
	migration.Register("Usercirce_20160906_204229", m)
}

// Run the migrations
func (m *Usercirce_20160906_204229) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("DROP TABLE IF EXISTS public.isla CASCADE;CREATE TABLE public.isla(id serial NOT NULL,ocupado bool NOT NULL,geometria geometry(POLYGON, 4326) NOT NULL,id_grupo_isla integer,CONSTRAINT isla_id_pkey PRIMARY KEY (id));COMMENT ON TABLE public.isla IS 'Es el espacio en que se pone la isla';COMMENT ON COLUMN public.isla.id IS 'Autoincremental para islas, ya que por ahora no hay otra forma de identificarlos';COMMENT ON COLUMN public.isla.ocupado IS 'Estado que indica si el espacio esta ocupado o no';COMMENT ON COLUMN public.isla.geometria IS 'El poligono que representa a la isla espacialmente';COMMENT ON COLUMN public.isla.id_grupo_isla IS 'Identificador del grupo isla';ALTER TABLE public.isla OWNER TO usercirce;ALTER TABLE public.isla ADD CONSTRAINT isla_grupo_isla_id_grupo_isla_fkey FOREIGN KEY (id_grupo_isla) REFERENCES public.grupo_isla (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
}

// Reverse the migrations
func (m *Usercirce_20160906_204229) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
