#!/bin/bash -ex

if [ "$TRAVIS" != true ]; then
  echo "This script is intended to run within the travis build only" 1>&2
  exit 1
fi

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o sqs-to-kafka-linux64
ln sqs-to-kafka-linux64 sqs-to-kafka
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o sqs-to-kafka-darwin64
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o sqs-to-kafka-windows64.exe

docker build --no-cache --pull --squash -t "meteogroup/sqs-to-kafka:$COMMIT" .
test/test.sh "$COMMIT"
