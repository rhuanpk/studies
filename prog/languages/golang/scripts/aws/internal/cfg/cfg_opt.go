package cfg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// Opt type fo handle with aws optFns.
type Opt func(*config.LoadOptions) error

// Creds struct to better handle with creds as params.
type Creds struct {
	AccessKey    string
	SecretKey    string
	SessionToken string
}

func loadCfg(ctx context.Context, opts ...Opt) (aws.Config, error) {
	var optFns []func(*config.LoadOptions) error

	for _, opt := range opts {
		optFns = append(optFns, opt)
	}

	return config.LoadDefaultConfig(ctx, optFns...)
}

// NewOpts get a new slice of [Opt] to get a new config.
func NewOpts(profile, region *string, creds *Creds) []Opt {
	var opts []Opt

	if profile != nil {
		opts = append(opts, Opt(config.WithSharedConfigProfile(*profile)))
	}

	if region != nil {
		opts = append(opts, Opt(config.WithRegion(*region)))
	}

	if creds != nil {
		opts = append(opts, Opt(config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				creds.AccessKey, creds.SecretKey, creds.SessionToken,
			),
		)))
	}

	return opts
}

// Get get a new config with empty context.
func Get(opts ...Opt) (aws.Config, error) {
	return loadCfg(context.Background(), opts...)
}

// GetWithContext get a new config with epecified context.
func GetWithContext(ctx context.Context, opts ...Opt) (aws.Config, error) {
	return loadCfg(ctx, opts...)
}
