package interceptor

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type AWSSTSProvider struct {
	RoleArn     string
	SessionName string
	Region      string
}

func (p *AWSSTSProvider) GetCredentials() (*Credentials, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(p.Region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := sts.NewFromConfig(cfg)
	input := &sts.AssumeRoleInput{
		RoleArn:         aws.String(p.RoleArn),
		RoleSessionName: aws.String(p.SessionName),
		DurationSeconds: aws.Int32(900), // 15 minutes
	}

	result, err := client.AssumeRole(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to assume role: %w", err)
	}

	creds := Credentials{
		AccessKeyID:     *result.Credentials.AccessKeyId,
		SecretAccessKey: *result.Credentials.SecretAccessKey,
		SessionToken:    *result.Credentials.SessionToken,
		Provider:        "AWS-STS",
	}
	return &creds, nil
}
