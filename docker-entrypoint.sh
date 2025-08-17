#!/bin/bash
set -e

mysql --ssl=0 -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -h "$MYSQL_HOST" -P "$MYSQL_PORT" -e "CREATE DATABASE IF NOT EXISTS $MYSQL_DATABASE;"

/usr/local/bin/migrate -path database/schema_migrate -database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_HOST:$MYSQL_PORT)/$MYSQL_DATABASE" up
echo "Schema Migrations Applied!"

if [[ "$SEED_DB" == "yes" ]]; then
    echo "Seeding dummy data..."
    mysql --ssl=0 -h "$MYSQL_HOST" -P "$MYSQL_PORT" -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" "$MYSQL_DATABASE" < database/dump.sql
else
    echo "Okay no probs... All the best."
fi

exec ./main
