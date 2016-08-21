#!/bin/bash
echo 'Ejecutando: install_utilities.sh'
list=(vim git nmap tree wget)
install=installed
for p in ${list[*]}
do
  if which $p&>/dev/null
  then
    echo "$p ya está instalado."
  else
    install=no_installed
  fi
done

if [ "$install" = "installed" ]
then
  echo 'Utilities ya están instalados. Nada que hacer.'
else
  sudo yum install -y ${list[*]}
fi
