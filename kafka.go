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
  "strings"
  kafka "github.com/Shopify/sarama"
)

type Producer kafka.SyncProducer

func createKafkaProducer() Producer {
  config := kafka.NewConfig()
  config.Producer.RequiredAcks = kafka.WaitForAll
  config.Producer.Retry.Max = 5
  config.Producer.Return.Successes = true

  brokers := strings.Split(kafkaBrokers, ",")
  producer, err := kafka.NewSyncProducer(brokers, config)
  logAndPanic(err)
  return producer
}

func publishMessage(message *string, producer Producer) (partition int32, offset int64, err error) {
  partition, offset, err = producer.SendMessage(&kafka.ProducerMessage{
    Topic: kafkaTopic,
    Value: kafka.StringEncoder(*message),
  })
  return
}
