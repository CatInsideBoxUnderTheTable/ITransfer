package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type FileAuthManager struct {
	ProfileName string
	Region      string
	session     *session.Options
}

func (f *FileAuthManager) OpenSession() {
	f.session = &session.Options{
		Profile: f.ProfileName,
		Config: aws.Config{
			Region: aws.String(f.Region),
		},
	}
}

func (f *FileAuthManager) GetExistingSession() any {
	return session.Must(session.NewSessionWithOptions(*f.session))
}

func (f *FileAuthManager) Close() {
	// todo read about session cleanup
}
