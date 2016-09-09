#!/bin/bash
sudo yum install -y httpd
sudo systemctl enable
sudo tee /etc/httpd/conf.d/permisos.conf << 'EOF'
User vagrant
Group vagrant
EOF
sudo tee /etc/httpd/conf.d/visor_parqueaderos.conf << 'EOF'
Alias "/visor" "/home/vagrant/src/Visor/app/"
<Directory "/home/vagrant/src/Visor/app/">
    Require all granted
</Directory>
ProxyPass /geoserver http://localhost:8000/geoserver
#ProxyPassReverse /geoserver http://localhost:8000/geoserver
EOF
sudo tee /etc/httpd/conf.d/cliente_parqueaderos.conf << 'EOF'
Alias "/parqueaderos" "/home/vagrant/src/Cliente/public"
<Directory "/home/vagrant/src/Cliente/public">
    Require all granted
</Directory>
ProxyPass /parqueaderos/v1 http://localhost:8080/parqueaderos/v1
EOF
