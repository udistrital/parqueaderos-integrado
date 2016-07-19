#!/bin/bash
function proxy_on {
export {http,HTTP,https,HTTPS,no,NO,ftp,FTP}_{proxy,PROXY}=http://10.20.4.15:3128
}
function proxy_off {
unset {http,HTTP,https,HTTPS,no,NO,ftp,FTP}_{proxy,PROXY}
}
proxy_on
