# Always specify version. Avoids bugs on image updates in hub
FROM mysql:8.0
#Copying db's initialization files
COPY ./sql-scripts/ /docker-entrypoint-initdb.d/

ENV MYSQL_ROOT_PASSWORD "root"

EXPOSE 3306 33060
