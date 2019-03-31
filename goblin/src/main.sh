#!/bin/bash

LC_ALL=en_US.UTF-8
LANGUAGE=en_US.UTF-8
# TODO load the tmp folder setting from env
TMP_DIR="/tmp/.goblin-pkg"
rm -rf $TMP_DIR
mkdir -p $TMP_DIR

sed -n -e '1,/^exit 0$/!p' $0 > "$TMP_DIR/bintop.tar.gz" 2>/dev/null
tar zxf $TMP_DIR/bintop.tar.gz -C $TMP_DIR >/dev/null 2>&1
rm -rf $TMP_DIR/bintop.tar.gz
mkdir -p /opt/pds/utils

cd ${TMP_DIR} && python main.py $@
rm -rf ${TMP_DIR}
exit 0
