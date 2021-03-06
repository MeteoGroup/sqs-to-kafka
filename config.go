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

import (
  "flag"
  "os"
)

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

func loadConfig() {
  flag.StringVar(&awsAccessKey, "aws-access-key", "", "AWS access key")
  flag.StringVar(&awsSecretKey, "aws-secret-key", "", "AWS secret key")
  flag.StringVar(&awsRegion, "aws-region", os.Getenv("AWS_REGION"), "AWS region")
  flag.StringVar(&awsProfile, "aws-profile", os.Getenv("AWS_PROFILE"), "AWS profile")
  flag.BoolVar(&awsReadConfig, "aws-read-config", false, "read AWS configuration from `~/.aws/config`")
  flag.StringVar(&awsEndpoint, "aws-endpoint", os.Getenv("AWS_ENDPOINT"), "URL of the AWS endpoint")
  flag.StringVar(&sqsUrl, "sqs-url", os.Getenv("SQS_URL"), "URL of the SQS queue for incomming messages")
  flag.StringVar(&kafkaBrokers, "kafka-brokers", os.Getenv("KAFKA_BROKERS"), "list of Kafka brokers used for bootstrapping")
  flag.StringVar(&kafkaTopic, "kafka-topic", os.Getenv("KAFKA_TOPIC"), "Kafka topic for outgoing messages")
  flag.StringVar(&metricsAddress, "metrics-address", os.Getenv("METRICS_ADDRESS"), "Listening address to serve metrics")
  flag.Parse()

  if sqsUrl == "" {
    panic("Required parameter `sqs-url` is missing.")
  }
  if kafkaBrokers == "" {
    panic("Required parameter `kafka-brokers` is missing.")
  }
  if kafkaTopic == "" {
    panic("Required parameter `kafka-topic` is missing.")
  }

  logConfig()
}

func logConfig() {
  parameters := []interface{}{}
  appendIfDefined := func(name string, value string) {
    if (value != "") {
      parameters = append(parameters, name, value)
    }
  }
  appendIfDefined("awsProfile", awsProfile)
  appendIfDefined("awsRegion", awsRegion)
  appendIfDefined("awsEndpoint", awsEndpoint)
  appendIfDefined("metricsAddress", metricsAddress)
  appendIfDefined("sqsUrl", sqsUrl)
  appendIfDefined("kafkaBrokers", kafkaBrokers)
  appendIfDefined("kafkaTopic", kafkaTopic)
  logInfo("starting with configuration", parameters...)
}
