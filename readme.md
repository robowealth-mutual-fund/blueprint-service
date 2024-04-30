# Blueprint golang service

## Require
- Go Version 19+ 

- Oracle db
  
```docker run -d -p 1521:1521 -e ORACLE_PASSWORD=PASSWORD APP_USER=TESTUSR APP_USER_PASSWORD=PASSWORD -v oracle-volume:/opt/oracle/oradata gvenzl/oracle-free```

https://github.com/gvenzl/oci-oracle-free#oracle-database-free-on-apple-m-chips

### Please change config on FIXME and read makefile