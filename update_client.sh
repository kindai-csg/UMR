#!/bin/bash

set -e

docker run -it --rm -v $PWD/client:/src -w /src node:13.8 yarn install
docker run -it --rm -v $PWD/client:/src -w /src node:13.8 yarn build