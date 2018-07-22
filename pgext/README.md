# A simple PG extension test

## Make
- make clean   
  Delete all object files and other build artifacts
- make all   
  Compile the extension
- make install    
  Install the extension
- make uninstall   
  Uninstall the extension by deleting the files
- make installcheck    
  Check if the extension is installed by running the tests. See:
  - ./sql/pg_simple_ext_test.sql - The test file
  - ./expected/pg_simple_ext.out - Expectations

## Resources
- https://www.postgresql.org/docs/current/static/xfunc-c.html
- http://big-elephants.com/2015-10/writing-postgres-extensions-part-i/
- http://big-elephants.com/2015-10/writing-postgres-extensions-part-ii/
- http://big-elephants.com/2015-10/writing-postgres-extensions-part-iii/
- http://big-elephants.com/2015-11/writing-postgres-extensions-part-iv/
- http://big-elephants.com/2015-11/writing-postgres-extensions-part-v/
- https://github.com/adjust/postgresql_extension_demo
- https://github.com/mrnugget/pg_hello_world

