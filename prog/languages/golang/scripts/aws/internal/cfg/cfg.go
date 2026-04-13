package cfg

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// GetConfig environment > ~/.aws/config and && ~/.aws/credentials
func GetConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalln("error in load default configs:", err)
	}
	return cfg
}

// GetConfigWithRegion like [GetConfig] but specify the region.
func GetConfigWithRegion(region string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalln("error in load default configs with region:", err)
	}
	return cfg
}

// GetConfigWithProfile like [GetConfig] but specify the profile.
func GetConfigWithProfile(profile string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithSharedConfigProfile(profile),
	)
	if err != nil {
		log.Fatalln("error in load default configs with profile:", err)
	}
	return cfg
}

// GetConfigFromStaticCreds get new config from specified credentials.
func GetConfigFromStaticCreds(accessKey, secretKey, sessionToken string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey, secretKey, sessionToken,
		)),
	)
	if err != nil {
		log.Fatalln("error in load default configs from static creds:", err)
	}
	return cfg
}
