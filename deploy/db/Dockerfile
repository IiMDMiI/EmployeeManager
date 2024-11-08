FROM postgres:14-alpine
COPY dumpfile.sql /docker-entrypoint-initdb.d/
COPY dbresetter.sh /docker-entrypoint-initdb.d/
COPY docker-entrypoint.sh usr/local/bin

#10 minutes for db reset
ENV RESET_DB_TIME 600 