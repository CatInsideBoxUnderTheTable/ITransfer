package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type FileAuthManager struct {
	ProfileName string
	Region      string
	session     *session.Options
}

func (this *FileAuthManager) OpenSession() {
	this.session = &session.Options{
		Profile: this.ProfileName,
		Config: aws.Config{
			Region: aws.String(this.Region),
		},
	}
}

func (this *FileAuthManager) GetExistingSession() any {
	return session.Must(session.NewSessionWithOptions(*this.session))
}

func (this *FileAuthManager) Close() {
	// todo read about session cleanup
}
