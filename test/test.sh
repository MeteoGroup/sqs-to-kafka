#!/bin/bash -e

cd "$(dirname "$0")"

TEST_STAGE_ID="$(printf '%04x'  "$RANDOM" "$RANDOM")"
TEST_STAGE="sqs2kafkatest$TEST_STAGE_ID"
export IMAGE_TAG="${1:-test}"

main() {
  trap 'tear-down' EXIT
  setup

  # when
  sqs send-message --queue-url http://sqs:9324/queue/messages --message-body "message text"

  # then
  local IFS=$'\n'
  messages_from_kafka=( $(read-kafka-messages) )

  assert "${#messages_from_kafka[@]}" == "1"
  assert ":${messages_from_kafka[0]}" == ":message text"
}

setup() {
  docker-compose -p "$TEST_STAGE" up -d sqs-to-kafka
}

tear-down() {
  result="$?"
  docker-compose -p "$TEST_STAGE" stop
  if (( result )); then
    docker-compose -p "$TEST_STAGE" logs
  fi
  docker-compose -p "$TEST_STAGE" down --volumes
  return "$result"
}

sqs() {
  in-test-stage -e 'AWS_ACCESS_KEY_ID=' -e 'AWS_SECRET_ACCESS_KEY=' \
    toolbelt/aws --region '(invalid)' --endpoint-url 'http://sqs:9324' sqs "$@"
}

read-kafka-messages() {
  in-test-stage ryane/kafkacat -q -b kafka:9092 -C -t sqs-messages -D $'\n' -e
}

in-test-stage() {
  docker run --rm --net "${TEST_STAGE}_default" "$@"
}

assert() {
  if ! test "$@"; then
    printf 'Assertion failed:'
    printf " '%s'" "$@"
    printf '\n'
    exit -1
  fi
} 1>&2

main "$@"
