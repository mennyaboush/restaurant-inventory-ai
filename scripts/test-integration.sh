#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR=$(cd "$(dirname "$0")/.." && pwd)
cd "$ROOT_DIR"

ENV_FILE=.env
EXAMPLE_FILE=.env.example

if [ ! -f "$ENV_FILE" ]; then
  if [ -f "$EXAMPLE_FILE" ]; then
    cp "$EXAMPLE_FILE" "$ENV_FILE"
    echo "[WARN] Created $ENV_FILE from $EXAMPLE_FILE. It contains placeholder credentials — replace with real secrets for CI."
  else
    echo "[WARN] No $EXAMPLE_FILE found to create $ENV_FILE; continuing with environment variables only."
  fi
fi

COMPOSE_YML=docker-compose.yml
USE_COMPOSE=0
if [ -f "$COMPOSE_YML" ]; then
  USE_COMPOSE=1
  echo "Using docker-compose to start services..."
  docker-compose up -d
else
  echo "docker-compose.yml not found; using fallback postgres container"
  CONTAINER_NAME=inventory-test-postgres
  if [ "$(docker ps -q -f name=$CONTAINER_NAME)" = "" ]; then
    # remove any stopped container with the same name
    if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" != "" ]; then
      docker rm -f $CONTAINER_NAME || true
    fi
    docker run --name $CONTAINER_NAME -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=testdb -p 5432:5432 -d postgres:15
  else
    echo "Postgres container $CONTAINER_NAME already running"
  fi
fi

# Read env file into current shell for the test runner
if [ -f "$ENV_FILE" ]; then
  # shellcheck disable=SC1090
  set -a; source "$ENV_FILE"; set +a
fi

HOST=${POSTGRES_HOST:-localhost}
PORT=${POSTGRES_PORT:-5432}
USER=${POSTGRES_USER:-postgres}

echo "Waiting for Postgres at $HOST:$PORT..."
MAX_WAIT=60
i=0
while ! (</dev/tcp/$HOST/$PORT) >/dev/null 2>&1; do
  i=$((i+1))
  if [ $i -ge $MAX_WAIT ]; then
    echo "Timed out waiting for Postgres at $HOST:$PORT"
    exit 2
  fi
  sleep 1
done

echo "Postgres is accepting connections. Running integration tests..."

TEST_ENV_FILE="$ENV_FILE" go test -tags=integration ./internal/repository -v
RC=$?

if [ $RC -eq 0 ]; then
  echo "Tests passed; tearing down test services..."
  if [ $USE_COMPOSE -eq 1 ]; then
    docker-compose down
  else
    docker stop inventory-test-postgres >/dev/null 2>&1 || true
    docker rm -f inventory-test-postgres >/dev/null 2>&1 || true
  fi
else
  echo "Tests failed (exit $RC). Leaving test services running for inspection." >&2
fi

exit $RC
