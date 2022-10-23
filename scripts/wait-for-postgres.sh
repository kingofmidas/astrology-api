#!/bin/sh
# wait-for-postgres.sh

host="$1"
shift
cmd="$@"
  
# Login for user (`-U`) and once logged in execute quit ( `-c \q` )
# If we can not login sleep for 1 sec
until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -U "$POSTGRES_USER" -c '\q'; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done

echo "Postgres is up - executing command"
# Print and execute all other arguments starting with `$1`
# So `exec "$1" "$2" "$3" ...`
exec "$@"