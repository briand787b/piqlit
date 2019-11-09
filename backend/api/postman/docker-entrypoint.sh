#!/bin/bash

/usr/local/bin/piqlit-cli block --timeout 5s
STATUS=$?
[ $STATUS -eq 0 ] || exit 1

exec newman run \
    --environment /etc/newman/environment.json \
    https://api.getpostman.com/collections/${POSTMAN_COLLECTION_ID}?apikey=${POSTMAN_API_KEY}