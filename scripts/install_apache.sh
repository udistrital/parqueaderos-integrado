#!/bin/bash
if httpd -v &> /dev/null
then
  echo 'HTTPD ya está instalado, nada que hacer.'
else
sudo yum install -y httpd
sudo systemctl enable httpd
fi

file=/etc/httpd/conf.d/permisos.conf
if [ -f $file ]
then
  echo "El archivo $file ya existe. Nada que hacer."
else
sudo tee $file << 'EOF'
User vagrant
Group vagrant
EOF
fi

file=/etc/httpd/conf.d/visor_parqueaderos.conf
if [ -f $file ]
then
  echo "El archivo $file ya existe. Nada que hacer."
else
sudo tee $file << 'EOF'
Alias "/visor" "/home/vagrant/src/Visor/app/"
<Directory "/home/vagrant/src/Visor/app/">
    Require all granted
</Directory>
ProxyPass /geoserver http://localhost:8000/geoserver
#ProxyPassReverse /geoserver http://localhost:8000/geoserver
EOF
fi

file=/etc/httpd/conf.d/cliente_parqueaderos.conf
if [ -f $file ]
then
  echo "El archivo $file ya existe. Nada que hacer."
else
sudo tee $file << 'EOF'
Alias "/parqueaderos" "/home/vagrant/src/Cliente/public"
<Directory "/home/vagrant/src/Cliente/public">
    Require all granted
</Directory>
ProxyPass /parqueaderos/v1 http://localhost:8080/parqueaderos/v1
EOF
fi

sudo systemctl start httpd
