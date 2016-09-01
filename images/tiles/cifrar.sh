#!/bin/bash
if [ "$1" != "" ]
then
    openssl aes-256-cbc -e -in $1 -out $1.AES256 &&
    rm -f $1
else
    for filename in $(ls | grep -v .sh)
    do
    openssl aes-256-cbc -e -in $filename -out $filename.AES256
    done
fi
