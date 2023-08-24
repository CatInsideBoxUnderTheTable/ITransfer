package main

import (
	"github.com/CatInsideBoxUnderTheTable/ITransfer/input"
	"github.com/CatInsideBoxUnderTheTable/ITransfer/storage"
	s3storage "github.com/CatInsideBoxUnderTheTable/ITransfer/storage/s3"
)

func main() {
	rawInput := input.GetUserConfig()

	s3Adapter := s3storage.S3Adapter{
		ConnectionManager: s3storage.FileAuthManager{
			Region:      rawInput.BucketRegion,
			ProfileName: rawInput.AuthFileProfile,
		},
	}
	tempUrl, err := uploadObject(rawInput, &s3Adapter)

	if err != nil {
		panic(err)
	}

	println(tempUrl)
}

func uploadObject(input input.UserInput, uploader storage.BucketUploader) (string, error) {
	uploader.InitializeSession(input.BucketName)
	uploader.PostObject(input.FilePath, input.FileName)

	result := uploader.GetObjectUrl(input.BucketName, input.ObjectLifeTimeInHours)

	return result, nil
}
