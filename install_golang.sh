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
mkdir -p /home/$USER/golang
sudo su -c "
echo 'export GOROOT=/usr/lib/golang
 export GOBIN=$GOROOT/bin
 export GOPATH=/home/$USER/golang
 export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' > /etc/profile.d/go.sh
"
echo '# Golang Path
 export GOROOT=/usr/lib/golang
 export GOBIN=$GOROOT/bin
 export GOPATH=/home/$USER/golang
 export PATH=$PATH:$GOROOT/bin$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
source /etc/profile
sudo su -c "
ldconfig
"
go version
go env
sudo su -c "
yum install -y git
"
go get github.com/gin-gonic/gin
