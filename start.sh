#!/bin/sh
set -e

echo "Stopping any running containers..."
docker compose down

echo "Building fresh images..."
docker compose build

echo "Starting database container..."
docker compose up -d db

echo "Waiting for the database to be ready..."
for i in $(seq 1 30); do
  if docker compose exec db mysqladmin ping -h"db" --silent; then
    echo "Database is ready!"
    break
  fi
  status=$(docker compose ps -q db | xargs docker inspect -f '{{.State.Status}}')
  if [ "$status" = "exited" ] || [ "$status" = "dead" ]; then
    echo "Database container failed to start. Showing last logs:"
    docker compose logs db
    exit 1
  fi
  sleep 2
done

if [ $i -eq 30 ]; then
  echo "Database did not become ready in time."
  docker compose logs db
  exit 1
fi

echo "Running migrations..."
if ! docker compose run --rm migrate; then
  echo "Migration failed. Showing logs:"
  docker compose logs migrate
  exit 1
fi

echo "Starting API and Webapp containers..."
docker compose up -d api webapp

echo "All services are up and running!"