#!/bin/bash
export LC_ALL=en_US.UTF-8
export LANGUAGE=en_US.UTF-8

TMP_DIR="/tmp/mess-tin"
rm -rf ${TMP_DIR}
mkdir -p ${TMP_DIR}

sed -n -e '1,/^exit 0$/!p' $0  | tar zxv -C ${TMP_DIR} 2>/dev/null
cd ${TMP_DIR} && python mess-tin.py $@

# remove folders (keep it for debug)
# rm -rf ${TMP_DIR}
exit 0
