FROM mysql:8.0.15
COPY ./init.sql /docker-entrypoint-initdb.d/
ENV MYSQL_ROOT_PASSWORD root
CMD ["mysqld"]
EXPOSE 3306
