FROM mysql:8.0.15
COPY ./init/* /docker-entrypoint-initdb.d/
ENV MYSQL_ROOT_PASSWORD root
CMD ["mysqld"]
