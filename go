#!/usr/bin/zsh
docker run --rm -it -p 8000:8000 -v "$PWD:/go/src/`basename $PWD`" -w /go/src/`basename $PWD` deis/go-dev:v0.23.0 go "$@"
