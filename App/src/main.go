package main

import (
	"github.com/aws/aws-sdk-go/service/s3/App/src/dataLayer/aws"
	domain "github.com/aws/aws-sdk-go/service/s3/App/src/domainLayer"
	presentation "github.com/aws/aws-sdk-go/service/s3/App/src/presentationLayer"
)

func main() {
	rawInput := presentation.GetUserConfig()

	adapter := aws.S3Adapter{
		ConnectionManager: &aws.FileAuthManager{
			Region:      *rawInput.UserEnvInput.BucketRegion,
			ProfileName: *rawInput.UserEnvInput.AuthFileProfile,
		},
	}
	domainInput := domain.UploadFileData{
		ObjectLifeTimeInHours: *rawInput.UserConsoleInput.ObjectLifeTimeInHours,
		FileName:              *rawInput.UserConsoleInput.FileName,
		FilePath:              *rawInput.UserConsoleInput.FilePath,
		BucketName:            *rawInput.UserEnvInput.BucketName,
	}

	proccessor := domain.RequestProcessor{Uploader: &adapter}
	var result = proccessor.UploadFileAndGenerateTemporaryLink(domainInput)

	println(result)
}
