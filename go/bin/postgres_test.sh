#!/bin/bash

docker-compose \
    -f ../docker-compose.yml \
    -f ../docker-compose.test.yml \
    down

docker-compose \
    -f ../docker-compose.yml \
    -f ../docker-compose.test.yml \
    build backend-test

docker-compose \
    -f ../docker-compose.yml \
    -f ../docker-compose.test.yml \
    run --rm backend-test go test github.com/briand787b/piqlit/core/postgres -race -run $1 -v

docker-compose \
    -f ../docker-compose.yml \
    -f ../docker-compose.test.yml \
    down