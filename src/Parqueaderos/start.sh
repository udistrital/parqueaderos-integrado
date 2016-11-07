#!/bin/bash
if [ ! -f Parqueaderos ]
then
  go build main.go
fi
cd /home/parking/parqueaderos/src/Parqueaderos
./Parqueaderos
