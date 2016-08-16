#!/bin/bash
wget http://apache.uniminuto.edu/maven/maven-3/3.3.9/binaries/apache-maven-3.3.9-bin.tar.gz
tar zxvf apache-maven-3.3.9-bin.tar.gz
sudo mv apache-maven-3.3.9 /usr/local
export PATH=/usr/local/apache-maven-3.3.9/bin:$PATH

