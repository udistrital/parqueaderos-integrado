insert into "tipo_propietario" ("nombre", "descripcion") values ('Algo', 'No se que es esto');
insert into "propietario" ("documento", "primer_nombre", "segundo_nombre", "primer_apellido", "segundo_apellido","id_tipo_propietario") values ('CC1024532159', 'Omar', 'Leonardo', 'Zambrano', 'Pulgarin','1');
insert into "veh" ("id_nfc", "placa", "id_propietario") values ('1234', 'ABC123', '3');
insert into "grupo_isla" ("geometria", "id_grupo_padre") values ('((6 15), (10 10), (20 10), (25 15), (25 35), (19 40), (11 40), (6 25), (6 15))', '1');
insert into "isla" ("estado", "id_vehiculo", "geometria", "hora_entrada", "hora_salida","id_grupo_isla") values ('1','1','((6,15), (10,10), (10,15))', '1999-01-08 04:05:06 -8:00','1999-01-08 05:05:06 -8:00','1');
