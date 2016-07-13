#!/bin/bash
sudo su -c "
yum –exclude=kernel* update -y
yum install -y gcc
yum install -y kernel-devel-$(uname -r)
yum install -y wget
wget -c http://download.virtualbox.org/virtualbox/5.0.20/VBoxGuestAdditions_5.0.20.iso -O VBoxGuestAdditions.iso
mount VBoxGuestAdditions.iso -o loop /mnt
/mnt/VBoxLinuxAdditions.run
#rm -rf VBoxGuestAdditions.iso
yum install -y golang
"
sudo su -c "
echo 'export GOROOT=/usr/lib/golang
 export GOBIN=$GOROOT/bin
 export GOPATH=/root
 export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' > /etc/profile.d/go.sh
"
sudo su vagrant -c "
echo '# Golang Path
 export GOROOT=/usr/lib/golang
 export GOBIN=$GOROOT/bin
 export GOPATH=/home/$USER
 export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
source /etc/profile
"
sudo su -c "
ldconfig
"
sudo su vagrant -c "
go version
go env
"
#go get github.com/gin-gonic/gin
