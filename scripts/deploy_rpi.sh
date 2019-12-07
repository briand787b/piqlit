#!/bin/bash

echo 'Make sure to execute this script from the root of project!'

docker login -u briand787b -p $DOCKER_HUB_PASSWORD

# build frontend-web-vue
docker image build -t briand787b/piqlit-vue-frontend:$(git rev-parse HEAD)-arm ./frontend/web/vue
docker image push briand787b/piqlit-vue-frontend:$(git rev-parse HEAD)-arm
docker image tag briand787b/piqlit-vue-frontend:$(git rev-parse HEAD)-arm briand787b/piqlit-vue-frontend:deployed-arm
docker image push briand787b/piqlit-vue-frontend:deployed-arm

# build postgresql
docker image build -t briand787b/piqlit-pg-db:$(git rev-parse HEAD)-arm ./backend/core/postgres
docker image push briand787b/piqlit-pg-db:$(git rev-parse HEAD)-arm
docker image tag briand787b/piqlit-pg-db:$(git rev-parse HEAD)-arm briand787b/piqlit-pg-db:deployed-arm
docker image push briand787b/piqlit-pg-db:deployed-arm

# build backend-go
docker image build \
    -t briand787b/piqlit-go-backend:$(git rev-parse HEAD)-arm \
    --build-arg arch=arm \
    --build-arg arm=7 \
    -f ./backend/api/rest/Dockerfile \
    ./backend
docker image push briand787b/piqlit-go-backend:$(git rev-parse HEAD)-arm
docker image tag briand787b/piqlit-go-backend:$(git rev-parse HEAD)-arm briand787b/piqlit-go-backend:deployed-arm
docker image push briand787b/piqlit-go-backend:deployed-arm