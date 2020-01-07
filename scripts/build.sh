#!/usr/bin/env bash

set -e


image() {
    docker build -t hello-world:latest -f build/docker/Dockerfile .
}

: ${1?}
$@
