package main

import (
	"github.com/aws/aws-sdk-go/service/s3/App/src/database/aws/adapters"
	"github.com/aws/aws-sdk-go/service/s3/App/src/database/aws/models"
)

func main() {
	adapter := adapters.AwsS3Adapter{}
	awsConnectionData := models.AwsConnectionData{
		Region:      "eu-central-1",
		ProfileName: "default",
	}
	adapter.InitializeSession("testenv-transferred-files-storage", awsConnectionData)
	adapter.PostObject("/home/pons/kotki.txt", "koteczki.txt")
	url := adapter.GetObjectUrl("koteczki.txt", 1)

	println(url)
}
