#!/bin/sh
set -e

echo "Starting database container..."
docker-compose up -d db

echo "Waiting for the database to be ready..."
until docker-compose exec db mysqladmin ping -h"db" --silent; do
  sleep 2
done

echo "Running migrations..."
docker-compose run --rm migrate

echo "Starting API and Webapp containers..."
docker-compose up -d api webapp
