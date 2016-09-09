#!/bin/bash
echo 'Ejecutando: install_nodejs.sh'
if npm --version&>/dev/null
then
  echo 'NodeJS y NPM ya están instalados. Nada que hacer.'
else

sudo curl --silent --location https://rpm.nodesource.com/setup_4.x | bash -
sudo yum install -y gcc-c++ make
sudo yum install -y nodejs
sudo yum install -y wget
wget https://npmjs.org/install.sh
sudo sh install.sh
npm --version

fi
