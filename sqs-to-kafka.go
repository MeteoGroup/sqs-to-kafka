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
  "runtime"
  "strconv"
)

func main() {
  logInfo("starting engines",
    "sqsUrl", sqsUrl,
    "awsRegion", awsRegion,
    "kafkaBrokers", kafkaBrokers,
    "kafkaTopic", kafkaTopic,
  )

  sqsClient := createSqsClient()
  kafkaProducer := createKafkaProducer()
  startPrometheusHttpExporter()

  logInfo("lift off")
  for {
    messages := fetchMessages(sqsClient)
    forwardedMessages, skippedMessages := forwardToKafka(messages, kafkaProducer)
    deleteMessageBatch(forwardedMessages, sqsClient)
    releaseMessageBatch(skippedMessages, sqsClient)
    runtime.GC()
  }
}

func forwardToKafka(messages Messages, kafkaProducer Producer) (forwarded Messages, skipped Messages) {
  forwarded = Messages{};
  skipped = Messages{};
  for _, message := range messages {
    partition, offset, err := publishMessage(message.Body, kafkaProducer)
    if (err != nil) {
      logError(err)
      skipped = append(skipped, message)
      messageCounter.WithLabelValues("skipped").Inc()
    } else {
      forwarded = append(forwarded, message)
      kafkaOffsets.WithLabelValues(strconv.FormatInt(int64(partition), 10)).Set(float64(offset))
      messageCounter.WithLabelValues("forwarded").Inc()
    }
  }
  return
}
