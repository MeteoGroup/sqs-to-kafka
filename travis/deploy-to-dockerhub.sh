#!/bin/sh

if [ "$TRAVIS" != true ]; then
  echo "This script is intended to run within the travis build only" 1>&2
  exit 1
fi

( # temporarily disable shell debugging, docker hub password would leak otherwise
  set +x
  docker login -u "$DOCKER_USER" -p "$DOCKER_PASS"
)

IMAGE_TAG="${TRAVIS_TAG:-"`date '+%Y%m%dT%H%M%S'`"}"
docker tag "meteogroup/sqs-to-kafka:$COMMIT" "meteogroup/sqs-to-kafka:$IMAGE_TAG"
docker tag "meteogroup/sqs-to-kafka:$COMMIT" "meteogroup/sqs-to-kafka:latest"
docker push "meteogroup/sqs-to-kafka:$IMAGE_TAG"
docker push "meteogroup/sqs-to-kafka:latest"
