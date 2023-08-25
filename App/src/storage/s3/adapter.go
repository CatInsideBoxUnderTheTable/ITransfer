package storage

import (
	"errors"
	"time"

	"github.com/CatInsideBoxUnderTheTable/ITransfer/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Adapter struct {
	ConnectionManager FileAuthManager
	bucketName        string
}

func (s *S3Adapter) InitializeSession(bucketName string) {
	s.ConnectionManager.OpenSession()
	s.bucketName = bucketName
}

func (s *S3Adapter) PostObject(filePath string, objectName string) error {
	file := utils.ReadFile(filePath)
	
    conn, err := convertToAwsSession(s.ConnectionManager.GetExistingSession())
    if err != nil{
        return err
    }

	uploader := s3manager.NewUploader(conn)
	uploaderInput := &s3manager.UploadInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(objectName),
		Body:   file,
	}

	var _, uploadErr = uploader.Upload(uploaderInput)

    defer utils.CloseFile(file)
    return uploadErr 
}

func (s *S3Adapter) GetObjectUrl(objectName string, urlExpirationHours uint) (string, error) {
	conn, err := convertToAwsSession(s.ConnectionManager.GetExistingSession())
    if err != nil{
        return "",err
    }

	client := s3.New(conn)
	downloaderInput := &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(objectName),
	}

	downloadRequest, _ := client.GetObjectRequest(downloaderInput)
	return createPresignedUrl(downloadRequest, urlExpirationHours)
}

func (s *S3Adapter) Close() {
	s.ConnectionManager.Close()
}

func convertToAwsSession(maybeSession any) (*session.Session, error) {
	awsSession, isAwsSession := maybeSession.(*session.Session)

	if !isAwsSession {
	    return nil, errors.New("Unable to resolve aws session")    
    }

	return awsSession, nil
}

func createPresignedUrl(downloadReq *request.Request, expirationTimeHours uint) (string, error) {
	duration := time.Hour * time.Duration(expirationTimeHours)
	downloadUrl, _, err := downloadReq.PresignRequest(duration)

	return downloadUrl, err
}
