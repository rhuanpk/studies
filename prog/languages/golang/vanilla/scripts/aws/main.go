package main

import (
	"dev/internal/cfg"
	s3Service "dev/internal/s3"
	sqsService "dev/internal/sqs"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func s3Execs() {
	const bucketName = "mybucket"
	s3Client := s3Service.GetClient(cfg.GetConfig())

	s3Service.ListObjects(s3Client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		// Prefix: aws.String("path"),
	})
	s3Service.ListObjectsPaginator(s3Client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		// Prefix: aws.String("path"),
	})
	s3Service.GetObject(s3Client, "file.any", &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("path/file.any"),
	})
	s3Service.DownloadObject(s3Client, "file.any", &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("path/file.any"),
	})
}

func sqsExecs() {
	const queueName = "myqueue"
	sqsClient := sqsService.GetClient(cfg.GetConfig())

	sqsService.ListQueues(sqsClient, &sqs.ListQueuesInput{
		QueueNamePrefix: aws.String(queueName),
	})
	sqsService.ListQueues(sqsClient, nil)
	sqsService.ListMessages(sqsClient, "queueURL")
	sqsService.DeleteMessage(sqsClient, "queueURL", "receiptHandle")
}

func main() {
	s3Execs()
	sqsExecs()
}
