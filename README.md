SQS to Kafka forwarder [![build status](https://travis-ci.org/MeteoGroup/sqs-to-kafka.svg)](https://travis-ci.org/MeteoGroup/sqs-to-kafka)
======================

Fetches messages from SQS and forwards them to Kafka


## Build

Given a properly setup go develop environment, just run

```
  go get -d -t .
  go build
```

The resulting binary should be named `sqs-to-kafka` (or `sqs-to-kafka.exe`
on Windows). For further detail, e.g. cross-platform builds and custom build
parameters, please refer to the go-documentation.


## Usage

`sqs-to-kafka` takes the following commandline parameters:

  - `--aws-access-key=STRING`: AWS access key
  - `--aws-secret-key=STRING`: AWS secret key
  - `--aws-region=STRING`: AWS region
  - `--aws-profile=STRING`: AWS profile
  - `--aws-read-config`: read AWS configuration from `~/.aws/config`
  - `--aws-endpoint=STRING`: URL of the AWS endpoint
  - `--sqs-url=STRING`: URL of the SQS queue for incomming messages
  - `--kafka-brokers=STRING`: list of Kafka brokers used for bootstrapping
  - `--kafka-topic=STRING`: Kafka topic for outgoing messages
  - `--metrics-address=HOST:PORT`: Listening address to serve metrics

`sqs-url`, `kafka-brokers` and `kafka-topic` are mandatory, everything else is
optional. When `metrics-address` is given `sqs-to-kafka` binds that address to
export prometheus compatible metrics.

In addition to that some AWS parameters may be passed as environment variables
as well, most notably:

  - `AWS_ACCESS_KEY_ID`: AWS access key
  - `AWS_SECRET_ACCESS_KEY`: AWS secret key
  - `AWS_REGION`: AWS region
  - `AWS_PROFILE`: AWS profile
  - `AWS_SDK_LOAD_CONFIG`: if set to `1` read AWS configuration from `~/.aws/config`

See the [AWS SDK documentation](https://docs.aws.amazon.com/sdk-for-go/api/aws/session/)
for a complete list. When both are specified, commandline parameters take
precedence over environment variables.


## Docker

A docker-ized variant is available `meteogroup/sqs-to-kafka`. Metrics are
exposed on port `8080`. To run use

```
docker run -P meteogroup/sqs-to-kafka <additional commandline arguments>
```


## License

Copyright © 2017 MeteoGroup Deutschland GmbH,
all the files in this repository are released under the terms of
[Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0).
