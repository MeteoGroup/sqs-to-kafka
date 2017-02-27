#!/bin/sh

if [ "$TRAVIS" != true ]; then
  echo "This script is intended to run within the travis build only" 1>&2
  exit 1
fi

docker tag "meteogroup/sqs-to-kafka:$COMMIT" "meteogroup/sqs-to-kafka:$IMAGE_TAG"
docker tag "meteogroup/sqs-to-kafka:$COMMIT" "meteogroup/sqs-to-kafka"
docker push "meteogroup/sqs-to-kafka:$IMAGE_TAG"
docker push "meteogroup/sqs-to-kafka"
