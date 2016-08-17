#!/bin/bash
echo 'Ejecutando: install_postgresql_postgis.sh'
sql1=usercirce.sql
sql2=script_tables.sql

if echo '\connect circe;' | sudo psql -U postgres&>/dev/null;then
  echo 'La base de datos ya está creada. Nada que hacer.'
else
  if echo '\connect circe; SELECT * FROM public.grupo_isla;' | sudo psql -U postgres&>/dev/null;then
    echo 'La base de datos ya está creada. Nada que hacer.'
  else

mv $sql1 /tmp
mv $sql2 /tmp
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
fi
