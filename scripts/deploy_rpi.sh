#!/bin/bash

echo 'Make sure to execute this script from the root of project!'

docker login -u briand787b -p $DOCKER_HUB_PASSWORD
STATUS=$?
[ $STATUS -eq 0 ] || exit 1

# build frontend-web-vue
echo Building Frontend...
docker image build -t briand787b/piqlit-vue-frontend:$(git rev-parse HEAD)-arm ./frontend/web/vue
docker image push briand787b/piqlit-vue-frontend:$(git rev-parse HEAD)-arm
docker image tag briand787b/piqlit-vue-frontend:$(git rev-parse HEAD)-arm briand787b/piqlit-vue-frontend:deployed-arm
docker image push briand787b/piqlit-vue-frontend:deployed-arm
echo Finished Building Frontend!

# build postgresql
echo Building Postgresql...
docker image build -t briand787b/piqlit-pg-db:$(git rev-parse HEAD)-arm ./backend/core/postgres
docker image push briand787b/piqlit-pg-db:$(git rev-parse HEAD)-arm
docker image tag briand787b/piqlit-pg-db:$(git rev-parse HEAD)-arm briand787b/piqlit-pg-db:deployed-arm
docker image push briand787b/piqlit-pg-db:deployed-arm
echo Finished Building Postgresql!

# build backend-go
echo Building Backend...
docker image build \
    -t briand787b/piqlit-go-backend:$(git rev-parse HEAD)-arm \
    --build-arg arch=arm \
    --build-arg arm=7 \
    -f ./backend/api/rest/Dockerfile \
    ./backend
docker image push briand787b/piqlit-go-backend:$(git rev-parse HEAD)-arm
docker image tag briand787b/piqlit-go-backend:$(git rev-parse HEAD)-arm briand787b/piqlit-go-backend:deployed-arm
docker image push briand787b/piqlit-go-backend:deployed-arm
echo Finished Building Backend!

read -r -p "deploy stack? [y/N] " response
if [[ ! "$response" =~ ^([yY][eE][sS]|[yY])+$ ]]
then
    exit 0
fi

echo Deploying Stack...
docker stack deploy -c docker-stack.arm.yml piqlit
echo Finished Deploying Stack!