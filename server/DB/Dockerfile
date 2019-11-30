FROM mysql:latest
COPY ./init/* /docker-entrypoint-initdb.d/
ENV MYSQL_ROOT_PASSWORD root
CMD ["mysqld"]
