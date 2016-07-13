#!/bin/bash
sudo su -c "
yum localinstall -y http://yum.postgresql.org/9.4/redhat/rhel-6-x86_64/pgdg-centos94-9.4-1.noarch.rpm
yum install -y postgresql94-server
/usr/pgsql-9.4/bin/postgresql94-setup initdb
systemctl enable postgresql-9.4
systemctl start postgresql-9.4.service
"
echo "Se ha terminado la instalación de Postgresql"
#if you what remove # yum erase postgresql94*
cat << 'EOF'
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
