#!/bin/bash

_NEXUS_SERVER=http://192.168.1.37:8101
_NPM_REGISTRY=${_NEXUS_SERVER}/repository/npm-host-home/
_MVN_REPO=mvnhome::::${_NEXUS_SERVER}/repository/mvn-home

echo "Removing dist folder"
rm -rvf ./dist 2>/dev/null
mkdir -pv ./dist/bundle 2>/dev/null

echo "downloading frontend package"
(cd ./dist/bundle && npm pack pasta@latest --registry=${_NPM_REGISTRY})
echo "downloading backend package"
mvn dependency:get -Ddest=./dist/bundle -DremoteRepositories=${_MVN_REPO} -Dartifact=xiongjia:portal:LATEST

echo "Creating tgz file"
(cd src && tar cvf ../dist/mess-tin.tar ./*.py)
(cd ./dist/bundle && tar rvf ../mess-tin.tar ./*)
gzip ./dist/mess-tin.tar
cat ./src/mess-tin.sh ./dist/mess-tin.tar.gz > ./dist/mess-tin.bin
