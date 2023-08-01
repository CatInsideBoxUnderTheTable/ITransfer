package main

import (
	aws "github.com/aws/aws-sdk-go/service/s3/App/src/dataLayer/aws"
	domain "github.com/aws/aws-sdk-go/service/s3/App/src/domainLayer"
	presentation "github.com/aws/aws-sdk-go/service/s3/App/src/presentationLayer"
)

func main() {
	adapter := aws.S3Adapter{
		ConnectionManager: &aws.FileAuthManager{
			Region:      "eu-central-1",
			ProfileName: "default",
		},
	}
	rawInput := presentation.ReadInputFromConsole()
	domainInput := domain.UploadFileData{
		ObjectLifeTimeInHours: *rawInput.ObjectLifeTimeInHours,
		FileName:              *rawInput.FileName,
		FilePath:              *rawInput.FilePath,
	}

	proccessor := domain.RequestProcessor{Uploader: &adapter}
	var result = proccessor.UploadFileAndGenerateTemporaryLink(domainInput)

	println(result)
}
