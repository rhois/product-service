#!/bin/sh
docker run --rm --network="host" -e DATABASE_URL=$1 -v ${PWD}/internal/infrastructure/db/psql/$2:/db/migrations amacneil/dbmate $3
