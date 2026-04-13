package sqs

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// GetClient get a new sqs client.
func GetClient(cfg aws.Config) *sqs.Client {
	return sqs.NewFromConfig(cfg)
}

// ListQueues list all sqs queues names.
func ListQueues(client sqs.ListQueuesAPIClient, params *sqs.ListQueuesInput) []string {
	var queueURLs []string
	paginator := sqs.NewListQueuesPaginator(client, params)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			log.Fatalln("error in get next sqs page")
		}

		log.Println("urls:", page.QueueUrls)
		queueURLs = append(queueURLs, page.QueueUrls...)
	}

	return queueURLs
}

// ListMessages list all messages in a sqs queue.
func ListMessages(client *sqs.Client, queueURL string) []types.Message {
	out, err := client.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: 10,
	})
	if err != nil {
		log.Fatalln("error in get sqs message:", err)
	}

	for _, message := range out.Messages {
		// log.Printf("message [%s]: %v", aws.ToString(message.ReceiptHandle), aws.ToString(message.Body))
		log.Printf("message [%s]: %v", aws.ToString(message.MessageId), aws.ToString(message.Body))
	}

	return out.Messages
}

// DeleteMessage deletes a message from sqs queue.
func DeleteMessage(client *sqs.Client, queueURL, receiptHandle string) {
	if _, err := client.DeleteMessage(context.Background(), &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	}); err != nil {
		log.Fatalln("error in delete sqs message:", err)
	}
}
