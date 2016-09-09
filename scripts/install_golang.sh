#!/bin/bash
echo 'Ejecutando: install_golang.sh'

if go version&>/dev/null
then
  echo 'Go ya está instalado. Nada que hacer.'
else

sudo yum install -y golang
sudo tee /etc/profile.d/go.sh << 'EOF'
export GOROOT=/usr/lib/golang
export GOBIN=$GOROOT/bin
export GOPATH=/root
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
EOF

file=/home/vagrant/.bashrc
tee $file << 'EOF'
# Golang Path
export GOROOT=/usr/lib/golang
export GOBIN=$GOROOT/bin
export GOPATH=/home/vagrant
export PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin
EOF
chown vagrant:vagrant $file

source ~/.bashrc
source /etc/profile
sudo ldconfig
go version
go env

fi
