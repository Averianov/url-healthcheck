FROM mysql:5.6

COPY ./deployments/mysql/create_db.sql /docker-entrypoint-initdb.d/create_db.sql

EXPOSE 3306
