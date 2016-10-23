#!/usr/bin/env bash

export DOCKER_REGISTRY=${DOCKER_REGISTRY:-""}
export DOCKER_IMAGE=${DOCKER_IMAGE:-"devlucky/fakelink"}
export DOCKER_TAG=${DOCKER_TAG:-"local"}

ENTRYPOINT=$1
ARGS=${@:2}

runDocker() {
  DOCKER_IMAGE=${DOCKER_IMAGE}-dev docker-compose run ${1} api ${ARGS}
}

if [ "$ENTRYPOINT" ]; then
  echo "Running docker with entrypoint $ENTRYPOINT $ARGS"
  runDocker "--entrypoint $ENTRYPOINT"
else
  echo "Running docker with default entrypoint and args"
  docker-compose up
fi

SUCCESS=$?

docker-compose kill
docker-compose rm -f

exit $SUCCESS