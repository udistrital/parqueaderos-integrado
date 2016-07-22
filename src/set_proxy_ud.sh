#!/bin/bash
if [ "$0" != "-bash" ]; then
  echo 'Ejecute el script con el comando $ source set_proxy_ud.sh'
  exit
fi
function proxy_on {
  export {http,HTTP,https,HTTPS,no,NO,ftp,FTP}_{proxy,PROXY}=http://10.20.4.15:3128
}
function proxy_off {
  unset {http,HTTP,https,HTTPS,no,NO,ftp,FTP}_{proxy,PROXY}
}
proxy_on
