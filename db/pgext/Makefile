#
# make clean        - Delete all object files and other build artifacts
# make all          - Compile the extension
# make install      - Install the extension
# make uninstall    - Uninstall the extension by deleting the files
# make installcheck - Check if the extension is installed by running the tests
# 										See:
#   									 > ./sql/pg_simple_ext_test.sql - The test file
#   									 > ./expected/pg_simple_ext.out - Expectations
#

EXTENSION = pg_simple_ext
MODULES   = pg_simple_ext
REGRESS   = pg_simple_ext_test
DATA      = pg_simple_ext--0.0.1.sql pg_simple_ext.control

PG_CONFIG = pg_config
PGXS := $(shell $(PG_CONFIG) --pgxs)
include $(PGXS)

