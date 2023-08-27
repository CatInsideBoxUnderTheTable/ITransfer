package storage

import "github.com/aws/aws-sdk-go/aws/credentials"

type awsPassProvider struct {
	accessKey       string
	secretAccessKey string
}

func (p *awsPassProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{
		ProviderName:    "ITransfer custom login provider",
		AccessKeyID:     p.accessKey,
		SecretAccessKey: p.secretAccessKey,
		SessionToken:    "", //no mfa enabled - no need for session token
	}, nil
}

// IsExpired never expires as this is hardcoded data
func (p *awsPassProvider) IsExpired() bool {
	return false
}

func (p *awsPassProvider) GetProvider() credentials.Provider {
	return p
}
