#!/bin/zsh
docker run --rm -it --net=host -v "$PWD:/app" -w /app ezynda3/php:7.1 php "$@"
