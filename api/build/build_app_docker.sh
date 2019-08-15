#!/usr/bin/env bash
set -ev

DOCKERFILE=Dockerfile

docker build -t ${PROJECT}:${TRAVIS_COMMIT} -f $DOCKERFILE .
