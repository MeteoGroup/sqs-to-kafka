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
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/service/sqs"
)

type Messages []*sqs.Message

func createSqsClient() *sqs.SQS {
  creds := (*credentials.Credentials)(nil)
  sharedConfig := session.SharedConfigStateFromEnv
  if awsReadConfig {
    sharedConfig = session.SharedConfigEnable
  }
  if (awsAccessKey != "" || awsSecretKey != "") {
    creds = credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, "")
  }
  awsSession, err := session.NewSessionWithOptions(session.Options{
    Config: aws.Config{
      Endpoint: optionalString(awsEndpoint),
      Region:   optionalString(awsRegion),
      Credentials: creds,
    },
    Profile: awsProfile,
    SharedConfigState: sharedConfig,
  })
  logAndPanic(err)
  sqsClient := sqs.New(awsSession)
  _, err = sqsClient.GetQueueAttributes(&sqs.GetQueueAttributesInput{QueueUrl: aws.String(sqsUrl)})
  logAndPanic(err)
  return sqsClient
}

func fetchMessages(sqsClient *sqs.SQS) Messages {
  response, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
    QueueUrl:            aws.String(sqsUrl),
    VisibilityTimeout:   aws.Int64(10),
    MaxNumberOfMessages: aws.Int64(10),
    WaitTimeSeconds:     aws.Int64(20),
  })
  if (err != nil) {
    logError(err)
    return Messages{};
  }
  messageCounter.WithLabelValues("received").Add(float64(len(response.Messages)))
  return response.Messages
}

func deleteMessageBatch(messages Messages, sqsClient *sqs.SQS) {
  if (len(messages) == 0) {
    return
  }
  toDelete := []*sqs.DeleteMessageBatchRequestEntry{}
  for _, deleted := range messages {
    toDelete = append(toDelete, &sqs.DeleteMessageBatchRequestEntry{
      Id:            deleted.MessageId,
      ReceiptHandle: deleted.ReceiptHandle,
    })
  }
  _, err := sqsClient.DeleteMessageBatch(&sqs.DeleteMessageBatchInput{
    QueueUrl: aws.String(sqsUrl),
    Entries:  toDelete,
  })
  logError(err)
}

func releaseMessageBatch(messages Messages, sqsClient *sqs.SQS) {
  if (len(messages) == 0) {
    return
  }
  toRelease := []*sqs.ChangeMessageVisibilityBatchRequestEntry{}
  for _, released := range messages {
    toRelease = append(toRelease, &sqs.ChangeMessageVisibilityBatchRequestEntry{
      Id: released.MessageId,
      ReceiptHandle: released.ReceiptHandle,
      VisibilityTimeout: aws.Int64(0),
    })
  }
  _, err := sqsClient.ChangeMessageVisibilityBatch(&sqs.ChangeMessageVisibilityBatchInput{
    QueueUrl: aws.String(sqsUrl),
    Entries: toRelease,
  })
  logError(err)
}

func optionalString(s string) *string {
  if (s == "") {
    return nil
  }
  return &s
}
