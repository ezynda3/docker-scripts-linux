#!/usr/bin/zsh
docker run --rm -it --net=host -p 8000:8000 -v "$PWD:/go/src/`basename $PWD`" -w /go/src/`basename $PWD` golang go "$@"
