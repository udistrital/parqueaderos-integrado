#!/bin/bash
echo 'Ejecutando: install_nodejs.sh'
if npm --version&>/dev/null; then
  echo 'NodeJS y NPM ya están instalados. Nada que hacer.'
else
sudo su -c "
curl --silent --location https://rpm.nodesource.com/setup_4.x | bash -
yum install -y gcc-c++ make
yum install -y nodejs
yum install -y wget
"
wget https://npmjs.org/install.sh
sudo su -c "
sh install.sh
"
npm --version
fi
