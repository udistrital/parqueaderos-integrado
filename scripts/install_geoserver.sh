#!/bin/bash
sudo yum install -y java-1.8.0-openjdk
#sudo yum install -y maven
sudo yum install -y unzip
wget http://pilotfiber.dl.sourceforge.net/project/geoserver/GeoServer/2.9.1/geoserver-2.9.1-bin.zip
unzip geoserver-*.zip
#cd geoserver-2.9.1
#mvn eclipse:eclipse -P wps
sudo cp geoserver-2.9.1 /usr/share/geoserver -R
echo "export GEOSERVER_HOME=/usr/share/geoserver" >> ~/.profile
. ~/.profile
sudo chown -R vagrant /usr/share/geoserver/
/usr/share/geoserver/bin/startup.sh

