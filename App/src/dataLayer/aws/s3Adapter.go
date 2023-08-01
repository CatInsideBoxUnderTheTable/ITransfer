package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	dataLayerAbstractions "github.com/aws/aws-sdk-go/service/s3/App/src/dataLayer/abstractions"
	"github.com/aws/aws-sdk-go/service/s3/App/src/utils"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"time"
)

type S3Adapter struct {
	ConnectionManager dataLayerAbstractions.AuthManager
	bucketName        string
}

func (this *S3Adapter) InitializeSession(bucketName string) {
	this.ConnectionManager.OpenSession()
	this.bucketName = bucketName
}

func (this *S3Adapter) PostObject(filePath string, objectName string) {
	file := utils.ReadFile(filePath)
	conn := convertAnyToAwsSession(this.ConnectionManager.GetExistingSession())

	uploader := s3manager.NewUploader(conn)
	uploaderInput := &s3manager.UploadInput{
		Bucket: aws.String(this.bucketName),
		Key:    aws.String(objectName),
		Body:   file,
	}

	var _, uploadErr = uploader.Upload(uploaderInput)
	utils.PanicIfErr(uploadErr, fmt.Sprintf("unable to upload %s to s3. INNER ERR: %s", objectName, uploadErr))

	defer utils.CloseFile(file)
}

func (this *S3Adapter) GetObjectUrl(objectName string, urlExpirationHours uint) string {
	conn := convertAnyToAwsSession(this.ConnectionManager.GetExistingSession())

	client := s3.New(conn)
	downloaderInput := &s3.GetObjectInput{
		Bucket: aws.String(this.bucketName),
		Key:    aws.String(objectName),
	}

	downloadRequest, _ := client.GetObjectRequest(downloaderInput)
	return createPresignedUrl(downloadRequest, urlExpirationHours)
}

func (this *S3Adapter) Close() {
	this.ConnectionManager.Close()
}

func convertAnyToAwsSession(maybeSession any) *session.Session {
	awsSession, isAwsSession := maybeSession.(*session.Session)

	if !isAwsSession {
		panic("Unable to resolve AWS session")
	}

	return awsSession
}

func createPresignedUrl(downloadRequest *request.Request, expiationTimeInHours uint) string {
	duration := time.Hour * time.Duration(expiationTimeInHours)
	downloadUrl, _, downloadUrlErr := downloadRequest.PresignRequest(duration)
	utils.PanicIfErr(downloadUrlErr, fmt.Sprintf("unable to presign object. INNER ERR: %s", downloadUrlErr))

	return downloadUrl
}
