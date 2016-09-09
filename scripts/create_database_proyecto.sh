#!/bin/bash
echo 'Ejecutando: create_database_proyecto.sh'

sql1=usercirce.sql
sql2=script_tables.sql
db=sigma4c

if echo 'SELECT * FROM public.grupo_isla;' | sudo psql -U postgres -d $db &>/dev/null
then
  echo 'La tabla grupo_isla ya está creada. Nada que hacer.'
elif echo "\connect $db;" | sudo psql -U postgres&>/dev/null
then
  echo 'La base de datos ya está creada. Nada que hacer.'
else
  cp $sql1 /tmp
  cp $sql2 /tmp
  scriptsql1=/tmp/$sql1
  scriptsql2=/tmp/$sql2
  sudo chown postgres:postgres $scriptsql1
  sudo chown postgres:postgres $scriptsql2
  sudo su postgres -c "
  cd /tmp
  psql -f $scriptsql1
  psql -f $scriptsql2 -d circe
  "
fi
