package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const queueName = "myqueue"

var ctx = context.Background()

func getClient() *sqs.Client {
	// ~/.aws/config and ~/.aws/credentials
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalln("error in load default configs:", err)
	}

	return sqs.NewFromConfig(cfg)
}

func listQueues(client sqs.ListQueuesAPIClient, params *sqs.ListQueuesInput) []string {
	var queueURLs []string
	paginator := sqs.NewListQueuesPaginator(client, params)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			log.Fatalln("error in get next sqs page")
		}

		log.Println("urls:", page.QueueUrls)
		queueURLs = append(queueURLs, page.QueueUrls...)
	}

	return queueURLs
}

func listMessages(client *sqs.Client, queueURL string) []types.Message {
	out, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
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

func deleteMessage(client *sqs.Client, queueURL, receiptHandle string) {
	if _, err := client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	}); err != nil {
		log.Fatalln("error in delete sqs message:", err)
	}
}

func main() {
	client := getClient()
	// listQueues(client, &sqs.ListQueuesInput{QueueNamePrefix: aws.String(queueName)})
	listQueues(client, nil)
	listMessages(client, "queueURL")
	deleteMessage(client, "queueURL", "receiptHandle")
}
