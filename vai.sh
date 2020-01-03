#!/bin/bash

set -euxo pipefail

cd ./test

docker-compose up --build
