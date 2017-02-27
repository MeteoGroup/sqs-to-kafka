#!/bin/sh -x

if [ "$TRAVIS" != true ]; then
  echo "This script is intended to run within the travis build only" 1>&2
  exit 1
fi

exec 3>&2 2>&- # disable logging to ensure we don't leak our docker secrets
docker login -e "$DOCKER_EMAIL" -u "$DOCKER_USER" -p "$DOCKER_PASS" 2>&3
exec 2>&3  3>&- # re-enable logging

docker tag "meteogroup/sqs-to-kafka:$COMMIT" "meteogroup/sqs-to-kafka:$IMAGE_TAG"
docker tag "meteogroup/sqs-to-kafka:$COMMIT" "meteogroup/sqs-to-kafka:latest"
docker push "meteogroup/sqs-to-kafka:$IMAGE_TAG"
docker push "meteogroup/sqs-to-kafka"
