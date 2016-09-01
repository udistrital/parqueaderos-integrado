#!/bin/bash
if [ "$1" != "" ]
then
    filenameout=$(echo $1 | sed 's/\.AES256$//g')
    openssl aes-256-cbc -d -in $1 -out $filenameout &&
    rm -f $1
else
    for filename in $(ls *.AES256)
    do
    filenameout=$(echo $filename | sed 's/\.AES256$//g')
    openssl aes-256-cbc -d -in $filename -out $filenameout
    done
fi
