package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type FileAuthManager struct {
	Login    string
	Password string
	Region   string
	session  *session.Options
}

func (f *FileAuthManager) OpenSession() {

	f.session = &session.Options{
		Config: aws.Config{
			Region:      aws.String(f.Region),
			Credentials: getAwsCredentials(f.Login, f.Password),
		},
	}
}

func (f *FileAuthManager) GetExistingSession() any {
	return session.Must(session.NewSessionWithOptions(*f.session))
}

func (f *FileAuthManager) Close() {
	// todo read about session cleanup
}

func getAwsCredentials(accessKey, secret string) *credentials.Credentials {
	init := awsPassProvider{
		secretAccessKey: secret,
		accessKey:       accessKey,
	}

	return credentials.NewCredentials(init.GetProvider())
}
