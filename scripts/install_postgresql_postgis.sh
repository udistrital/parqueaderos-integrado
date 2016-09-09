#!/bin/bash
echo 'Ejecutando: install_postgresql_postgis.sh'
if [ -f /usr/pgsql-9.4/bin/pg_ctl ]; then
  echo 'Postgres ya está instalado. Nada que hacer.'
else
echo "Instalando PostgreSQL"
sudo yum localinstall -y http://yum.postgresql.org/9.4/redhat/rhel-6-x86_64/pgdg-centos94-9.4-1.noarch.rpm
sudo yum install -y postgresql94-server
sudo /usr/pgsql-9.4/bin/postgresql94-setup initdb
sudo systemctl enable postgresql-9.4
sudo sed -i.bak 's/peer/trust/; s/ident/md5/' /var/lib/pgsql/9.4/data/pg_hba.conf
sudo tee -a /var/lib/pgsql/9.4/data/pg_hba.conf << 'EOF'
host    all             all             0.0.0.0/0               md5
EOF
sudo sed -i.bak \"s/#listen_addresses = 'localhost'/listen_addresses = '*'/\" /var/lib/pgsql/9.4/data/postgresql.conf
sudo systemctl start postgresql-9.4.service
echo "Se ha terminado la instalación de Postgresql"
#if you what remove # yum erase postgresql94*
cat << 'EOF'
# ¿Cómo crear un usuario y una base de datos?
$ sudo su postgres
$ psql
> CREATE USER nombre_usuario WITH PASSWORD 'clave';
> CREATE DATABASE nombre_db WITH OWNER = nombre_usuario ENCODING = 'UTF8' TABLESPACE = pg_default CONNECTION LIMIT = -1;
$ vim /var/lib/pgsql/9.4/data/pg_hba.conf
#host   DATABASE        USER            ADDRESS                 METHOD
local   all             all                                     trust
host    all             all             127.0.0.1/32            md5
host    all             all             localhost               md5
host    all             all             0.0.0.0/0               md5
host    all             all             ::1/128                 md5#sara
$ sudo systemctl restart postgresql-9.4.service
EOF
echo "Instalando Postgis"
sudo yum install -y epel-release
sudo yum install -y postgis2_94
echo "Postgis Instalado"
cat << 'EOF'
# ¿Cómo instalar postgis en tu base de datos?
$ sudo su postgres
$ /usr/pgsql-9.4/bin/psql -p 5432
> CREATE DATABASE gistest;
> \connect gistest;
> CREATE EXTENSION postgis;
> CREATE EXTENSION postgis_topology;
> CREATE EXTENSION ogr_fdw;
> SELECT postgis_full_version();
EOF
fi
