#!/bin/sh
set -e

echo "Waiting for PostgreSQL to be ready..."
echo "postgres:5432:${POSTGRES_DB}:${POSTGRES_USER}:${POSTGRES_PASSWORD}" >~/.pgpass
chmod 600 ~/.pgpass
export PGPASSFILE=~/.pgpass

until PGPASSWORD="$POSTGRES_PASSWORD" psql -h postgres -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c '\q'; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 2
done

echo "Running database migrations..."
task migrate:up || echo "Migration failed!"
echo "Database migrations complete."

# Start cron in background
crond

# Make sure the pre-built binary has proper permissions

if [ ! -f /app/tmp/server ]; then
  echo "Binary not found, building application..."
  mkdir -p /app/tmp
  go build -o /app/tmp/server .
fi

chmod +x /app/tmp/server || echo "Could not set executable permission, continuing anyway"

# Choose whether to use Air for development or direct execution
if [ "${USE_AIR:-true}" = "true" ]; then
  echo "Starting Air with pre-built binary..."
  exec air -c /app/.air.toml
else
  echo "Starting the server directly..."
  exec /app/tmp/server
fi
