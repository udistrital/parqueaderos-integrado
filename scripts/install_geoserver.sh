#!/bin/bash
echo 'Ejecutando: install_geoserver.sh'

if [ -f /usr/share/geoserver/bin/startup.sh ]
then
  echo 'GeoServer ya está instalado. Nada que hacer.'
else

sudo yum install -y java-1.8.0-openjdk
#sudo yum install -y maven
sudo yum install -y unzip
echo 'Descargando GeoServer... Espere un momento'
wget http://pilotfiber.dl.sourceforge.net/project/geoserver/GeoServer/2.9.1/geoserver-2.9.1-bin.zip > /dev/null 2>&1
unzip geoserver-*.zip
#cd geoserver-2.9.1
#mvn eclipse:eclipse -P wps
sudo cp geoserver-2.9.1 /usr/share/geoserver -R
echo "export GEOSERVER_HOME=/usr/share/geoserver" >> ~/.profile
. ~/.profile
sudo chown -R vagrant /usr/share/geoserver/
#/usr/share/geoserver/bin/startup.sh &
sudo tee /usr/lib/systemd/system/geoserver.service << 'EOF'
[Unit]
Description=geoserver
After=network.target
#After=network.target remote-fs.target nss-lookup.target

[Service]
Environment=GEOSERVER_HOME=/usr/share/geoserver
ExecStart=/usr/share/geoserver/bin/startup.sh
Type=simple
#ExecStop=/usr/lib/systemd/scripts/apachectl stop
#RemainAfterExit=yes

[Install]
WantedBy=default.target
EOF
sudo sed -i.bak 's/jetty.port=8080/jetty.port=8000/g' /usr/share/geoserver/start.ini
sudo systemctl enable geoserver
sudo systemctl start geoserver

fi
