#!/bin/bash

docker run \
    -d \
    -p 5432:5432 \
    -e POSTGRES_PASSWORD=s3cr3t \
    postgres:14-alpine
