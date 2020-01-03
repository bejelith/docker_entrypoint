#!/bin/bash

set -euxo pipefail

cd ./test

docker-compose up --build --abort-on-container-exit --scale=docker_entrypoint_integration=0
docker-compose up --build --exit-code-from docker_entrypoint_integration --scale docker_entrypoint_tests=0
