/*
Copyright © 2017 MeteoGroup Deutschland GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
 */
package main

import "flag"

var (
  awsProfile = ""
  awsAccessKey = ""
  awsSecretKey = ""
  awsRegion = ""
  sqsUrl = ""
  awsEndpoint = ""
  kafkaBrokers = ""
  kafkaTopic = ""
  metricsAddress = ""
  awsReadConfig = false
)

func init() {
  flag.StringVar(&awsAccessKey, "aws-access-key", "", "AWS access key")
  flag.StringVar(&awsSecretKey, "aws-secret-key", "", "AWS secret key")
  flag.StringVar(&awsRegion, "aws-region", "", "AWS region")
  flag.StringVar(&awsProfile, "aws-profile", "", "AWS profile")
  flag.BoolVar(&awsReadConfig, "aws-read-config", false, "read AWS configuration from `~/.aws/config`")
  flag.StringVar(&awsEndpoint, "aws-endpoint", "", "URL of the AWS endpoint")
  flag.StringVar(&sqsUrl, "sqs-url", "", "URL of the SQS queue for incomming messages")
  flag.StringVar(&kafkaBrokers, "kafka-brokers", "", "list of Kafka brokers used for bootstrapping")
  flag.StringVar(&kafkaTopic, "kafka-topic", "", "Kafka topic for outgoing messages")
  flag.StringVar(&metricsAddress, "metrics-address", "", "Listening address to serve metrics")
  flag.Parse()

  if sqsUrl == "" {
    panic("Required parameter `--sqs-url` is missing.")
  }
  if kafkaBrokers == "" {
    panic("Required parameter `--kafka-brokers` is missing.")
  }
  if kafkaTopic == "" {
    panic("Required parameter `--kafka-topic` is missing.")
  }
}
