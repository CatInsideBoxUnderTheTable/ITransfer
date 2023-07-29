package adapters

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/App/src/database/aws/models"
	"github.com/aws/aws-sdk-go/service/s3/App/src/utils"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"time"
)

type AwsS3Adapter struct {
	awsSession *session.Session
	bucketName string
}

func (this *AwsS3Adapter) InitializeSession(bucketName string, connectionData models.AwsConnectionData) {
	sessionOptions := *prepareConnectionConfig(connectionData)
	this.awsSession = session.Must(session.NewSessionWithOptions(sessionOptions))
	this.bucketName = bucketName
}

func (this *AwsS3Adapter) PostObject(fullPath string, objectName string) {
	file := readFile(fullPath)
	uploader := s3manager.NewUploader(this.awsSession)
	uploaderInput := &s3manager.UploadInput{
		Bucket: aws.String(this.bucketName),
		Key:    aws.String(objectName),
		Body:   file,
	}

	var _, uploadErr = uploader.Upload(uploaderInput)
	utils.PanicIfErr(uploadErr, fmt.Sprintf("unable to upload %s to s3. INNER ERR: %s", objectName, uploadErr))
}

func (this *AwsS3Adapter) GetObjectUrl(objectName string, urlExpirationHours uint) string {
	client := s3.New(this.awsSession)
	downloaderInput := &s3.GetObjectInput{
		Bucket: aws.String(this.bucketName),
		Key:    aws.String(objectName),
	}

	downloadRequest, _ := client.GetObjectRequest(downloaderInput)

	return createPresignedUrl(downloadRequest, urlExpirationHours)
}

func prepareConnectionConfig(connectionData models.AwsConnectionData) *session.Options {
	return &session.Options{
		Profile: connectionData.ProfileName,
		Config: aws.Config{
			Region: aws.String(connectionData.Region),
		},
	}
}

func createPresignedUrl(downloadRequest *request.Request, expiationTimeInHours uint) string {
	duration := time.Hour * time.Duration(expiationTimeInHours)

	downloadUrl, _, downloadUrlErr := downloadRequest.PresignRequest(duration)
	utils.PanicIfErr(downloadUrlErr, fmt.Sprintf("unable to presign object. INNER ERR: %s", downloadUrlErr))

	return downloadUrl
}

func readFile(path string) *os.File {
	file, fileErr := os.Open(path)
	utils.PanicIfErr(fileErr, fmt.Sprintf("file at path %s cannot be opened", path))

	return file
}
