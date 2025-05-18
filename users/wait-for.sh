#!/bin/sh
set -e

host="$1"
shift
cmd="$@"

until pg_isready -h "$host" -p 5432 > /dev/null 2>&1; do
  echo "Esperando a Postgres en $host:5432..."
  sleep 1
done

echo "✅ Postgres está listo, ejecutando: $cmd"
exec $cmd
