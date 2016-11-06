#!/bin/bash
sudo tee usr/lib/systemd/system/parqueaderos.service << EOF
[Unit]
Description=parqueaderos
After=network.target
#After=network.target remote-fs.target nss-lookup.target

[Service]
#Environment=GEOSERVER_HOME=/usr/share/geoserver
ExecStart=$GOPATH/bin/Parqueaderos
Type=simple
#ExecStop=/usr/lib/systemd/scripts/apachectl stop
#RemainAfterExit=yes

[Install]
WantedBy=default.target
EOF
