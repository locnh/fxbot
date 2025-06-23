#!/bin/bash
echo "$DOCKER_PASSWORD" | docker login --username "$DOCKER_USERNAME" --password-stdin

if [ $TRAVIS_BRANCH != "master" ]; then
    docker tag locnh/fxbot locnh/fxbot:$TRAVIS_BRANCH
    docker push locnh/fxbot:$TRAVIS_BRANCH
else
    docker push locnh/fxbot
fi
