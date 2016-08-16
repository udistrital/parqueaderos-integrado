#!/bin/bash
git clone git://github.com/geoserver/geoserver.git geoserver --depth=1 -b master
cd geoserver/src
mvn clean install
mvn eclipse:eclipse
#needs maven  3.1
