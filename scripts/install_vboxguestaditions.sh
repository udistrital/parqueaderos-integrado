#!/bin/bash
echo 'Ejecutando: install_vboxguestaditions.sh'
if lsmod | grep -i vboxguest &> /dev/null; then
 echo 'VirtualBox Guest Additions ya está instalado. Nada que hacer.'
else

sudo yum --exclude=kernel* update -y
sudo yum install -y gcc
sudo yum install -y kernel-devel-$(uname -r)
sudo yum install -y wget
echo 'Descargando VBox image.. Espere un momento.'
wget -c http://download.virtualbox.org/virtualbox/5.0.20/VBoxGuestAdditions_5.0.20.iso -O VBoxGuestAdditions.iso > /dev/null 2>&1
echo 'Descarga Terminada.'
sudo mount VBoxGuestAdditions.iso -o loop /mnt
sudo /mnt/VBoxLinuxAdditions.run
#rm -rf VBoxGuestAdditions.iso

fi
