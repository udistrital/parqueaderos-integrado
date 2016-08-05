#!/bin/bash
file=circe.sql
postgreshome=/var/lib/pgsql

scriptsql=$postgreshome/$file
sudo chown postgres:postgres $scriptsql
sudo su postgres -c "
psql -f $scriptsql
"

