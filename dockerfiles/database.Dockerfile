# Use the official PostgreSQL image as the base image
FROM postgres:17.4-alpine3.21

# Create a directory to store the PostgreSQL data

# Set environment variables for the database
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=dbname
ENV PGDATA /var/lib/postgresql/data

VOLUME $PGDATA

EXPOSE 5432

CMD ["postgres"]


