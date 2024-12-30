FROM postgres:latest
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=gozime32
ENV POSTGRES_DB=XO
COPY iternal/database/sql/create_tables.sql /docker-entrypoint-initdb.d/
