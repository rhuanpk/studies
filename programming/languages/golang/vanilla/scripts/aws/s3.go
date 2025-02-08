package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const bucketName = "mybucket"

var ctx = context.Background()

func getClient() *s3.Client {
	// ~/.aws/config and ~/.aws/credentials
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalln("error in load default configs:", err)
	}

	return s3.NewFromConfig(cfg)
}

func listObjects(client *s3.Client, params *s3.ListObjectsV2Input) []string {
	var objNames []string

	out, err := client.ListObjectsV2(ctx, params)
	if err != nil {
		log.Fatalln("error in list s3 objects:", err)
	}

	for _, content := range out.Contents {
		log.Println("key:", aws.ToString(content.Key))
		objNames = append(objNames, aws.ToString(content.Key))
	}

	return objNames
}

func listObjectsPaginator(client *s3.Client, params *s3.ListObjectsV2Input) {
	paginator := s3.NewListObjectsV2Paginator(client, params)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			log.Fatalln("error in get next bucket page:", err)
		}

		for _, content := range page.Contents {
			log.Println("key:", aws.ToString(content.Key))
		}
	}
}

func getObject(fileName string, client *s3.Client, params *s3.GetObjectInput) {
	obj, err := client.GetObject(ctx, params)
	if err != nil {
		log.Fatalln("error in get s3 object:", err)
	}
	defer obj.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("error in os create file:", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, obj.Body); err != nil {
		log.Fatalln("error in copy body to file:", err)
	}
}

func downloadObject(fileName string, client manager.DownloadAPIClient, input *s3.GetObjectInput) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("error in os create file:", err)
	}
	defer file.Close()

	if _, err := manager.NewDownloader(client).Download(ctx, file, input); err != nil {
		log.Fatalln("error in dowload s3 object:", err)
	}
}

func main() {
	client := getClient()

	listObjects(client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		// Prefix: aws.String("path"),
	})

	listObjectsPaginator(client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		// Prefix: aws.String("path"),
	})

	getObject("file.any", client, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("path/file.any"),
	})

	downloadObject("file.any", client, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("path/file.any"),
	})
}
