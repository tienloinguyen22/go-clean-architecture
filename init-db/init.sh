#!/bin/bash
set -e

# Start PostgreSQL in the background
docker-entrypoint.sh postgres &

# Wait until it's ready
until pg_isready -U "$POSTGRES_USER"; do
  echo "Waiting for postgres..."
  sleep 1
done

# Check and create the DB if needed
DB_EXISTS=$(psql -U "$POSTGRES_USER" -tAc "SELECT 1 FROM pg_database WHERE datname='go-clean-architecture'")
if [ "$DB_EXISTS" != "1" ]; then
  echo "Creating database 'go-clean-architecture'..."
  createdb -U "$POSTGRES_USER" go-clean-architecture
  echo "Database 'go-clean-architecture' created successfully."
else
  echo "Database 'go-clean-architecture' already exists."
fi

# Bring foreground process back
wait
