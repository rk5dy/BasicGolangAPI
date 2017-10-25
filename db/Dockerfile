FROM postgres:10

ADD firecontrol.sql /docker-entrypoint-initdb.d/

RUN chmod 755 /docker-entrypoint-initdb.d/*.sql

EXPOSE 5432

HEALTHCHECK --interval=10s --timeout=5s --retries=6 --start-period=30s CMD timeout --preserve-status 5 nc -z -w 5 formation-database $DB_PORT
