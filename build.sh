#!/usr/bin/env bash
set -e

DOCKER_VERSION=v1

## Server

DOCKER_SERVER=lucasponce/server
DOCKER_SERVER_TAG=${DOCKER_SERVER}:${DOCKER_VERSION}

rm -Rf docker/server/server
cd server
go build -o ../docker/server/server
cd ..

podman build -t ${DOCKER_SERVER_TAG} docker/server

## Client

DOCKER_CLIENT=lucasponce/client
DOCKER_CLIENT_TAG=${DOCKER_CLIENT}:${DOCKER_VERSION}

rm -Rf docker/client/client
cd client
go build -o ../docker/client/client
cd ..

podman build -t ${DOCKER_CLIENT_TAG} docker/client


podman login docker.io
podman push ${DOCKER_SERVER_TAG}
podman push ${DOCKER_CLIENT_TAG}
