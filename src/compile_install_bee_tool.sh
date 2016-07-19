#!/bin/bash
sudo su vagrant -c "
cd /home/vagrant/src/github.com/beego/bee/
go build
"
sudo cp /home/vagrant/src/github.com/beego/bee/bee /usr/lib/golang/bin/
