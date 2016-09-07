package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Usercirce_20160906_205531 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Usercirce_20160906_205531{}
	m.Created = "20160906_205531"
	migration.Register("Usercirce_20160906_205531", m)
}

// Run the migrations
func (m *Usercirce_20160906_205531) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("DROP TABLE IF EXISTS public.registro CASCADE;CREATE TABLE public.registro(id serial NOT NULL,id_vehiculo integer NOT NULL,id_isla integer NOT NULL,hora_entrada timestamp,hora_salida timestamp,CONSTRAINT registro_pkey PRIMARY KEY (id));COMMENT ON TABLE public.registro IS 'Es el espacio donde queda almacenado el registro de entrada y salida de un vehiculo al parqueadero';COMMENT ON COLUMN public.registro.id IS 'Autoincremental para registros, ya que por ahora no hay otra forma de identificarlos';COMMENT ON COLUMN public.registro.id_vehiculo IS 'identificador del vehiculo y del propietario';COMMENT ON COLUMN public.registro.id_isla IS 'identificador de la isla';COMMENT ON COLUMN public.registro.hora_entrada IS 'Hora en que ingresa el vehiculo';COMMENT ON COLUMN public.registro.hora_salida IS 'Hora en que sale el vehiculo';ALTER TABLE public.registro OWNER TO usercirce;ALTER TABLE public.registro ADD CONSTRAINT registro_vehiculo_id_vehiculo_fkey FOREIGN KEY (id_vehiculo) REFERENCES public.vehiculo (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;ALTER TABLE public.registro ADD CONSTRAINT registro_isla_id_isla_fkey FOREIGN KEY (id_isla) REFERENCES public.isla (id) MATCH FULL ON DELETE NO ACTION ON UPDATE NO ACTION;")
}

// Reverse the migrations
func (m *Usercirce_20160906_205531) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
