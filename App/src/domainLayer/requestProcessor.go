package domain

import dataLayerAbstractions "github.com/aws/aws-sdk-go/service/s3/App/src/dataLayer/abstractions"

type RequestProcessor struct {
	Uploader dataLayerAbstractions.BucketUploader
}

func (this *RequestProcessor) UploadFileAndGenerateTemporaryLink(requestData UploadFileData) string {
	this.Uploader.InitializeSession(requestData.BucketName)
	this.Uploader.PostObject(requestData.FilePath, requestData.FileName)

	return this.Uploader.GetObjectUrl(requestData.FileName, requestData.ObjectLifeTimeInHours)
}
