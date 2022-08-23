## Using the needed mysql docker version
FROM mysql:5.7


ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE: g_wc_ex_borrar

ADD init.sql /docker-entrypoint-initdb.d/init.sql

## Changing default shell to bash
SHELL ["/bin/bash", "-c"]

EXPOSE 3306