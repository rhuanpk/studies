package s3

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// GetClient get a new s3 client.
func GetClient(cfg aws.Config) *s3.Client {
	return s3.NewFromConfig(cfg)
}

// ListObjects list the s3 objects.
func ListObjects(client *s3.Client, params *s3.ListObjectsV2Input) []string {
	var objNames []string

	out, err := client.ListObjectsV2(context.Background(), params)
	if err != nil {
		log.Fatalln("error in list s3 objects:", err)
	}

	for _, content := range out.Contents {
		log.Println("key:", aws.ToString(content.Key))
		objNames = append(objNames, aws.ToString(content.Key))
	}

	return objNames
}

// ListObjectsPaginator list the s3 objects from paginator strategy.
func ListObjectsPaginator(client *s3.Client, params *s3.ListObjectsV2Input) {
	paginator := s3.NewListObjectsV2Paginator(client, params)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.Background())
		if err != nil {
			log.Fatalln("error in get next bucket page:", err)
		}

		for _, content := range page.Contents {
			log.Println("key:", aws.ToString(content.Key))
		}
	}
}

// GetObject get a single object from s3.
func GetObject(client *s3.Client, fileName string, params *s3.GetObjectInput) {
	obj, err := client.GetObject(context.Background(), params)
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

// DownloadObject downloads a s3 object to the local machine.
func DownloadObject(client manager.DownloadAPIClient, fileName string, input *s3.GetObjectInput) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("error in os create file:", err)
	}
	defer file.Close()

	if _, err := manager.NewDownloader(client).Download(context.Background(), file, input); err != nil {
		log.Fatalln("error in dowload s3 object:", err)
	}
}
