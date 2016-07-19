#!/bin/bash

repobee="github.com/beego/bee"

archivos=( 
"g_appcode.go"
)

dependencias=(
"github.com/go-sql-driver/mysql"
"github.com/lib/pq"
"github.com/astaxie/beego"
)

mkdir /tmp/bee
for i in "${archivos[@]}"; do
  mv /home/vagrant/src/$repobee/$i /tmp/bee/$i
done

rm -rf /home/vagrant/src/github.com
for i in "${dependencias[@]}"; do
  rm -rf /home/vagrant/src/$i
  go get $i
done

rm -rf /home/vagrant/src/$repobee
go get $repobee

cd /home/vagrant/src/github.com/beego/bee/
for i in "${archivos[@]}"; do
  cp -f /tmp/bee/$i $i
done
go build

sudo cp /home/vagrant/src/github.com/beego/bee/bee /usr/lib/golang/bin/
