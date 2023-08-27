package main

import (
	"github.com/CatInsideBoxUnderTheTable/ITransfer/input"
	"github.com/CatInsideBoxUnderTheTable/ITransfer/storage"
	s3storage "github.com/CatInsideBoxUnderTheTable/ITransfer/storage/s3"
)

func main() {
	rawInput, err := input.GetUserConfig()
	if err != nil {
		panic(err)
	}

	s3Adapter := s3storage.S3Adapter{
		ConnectionManager: s3storage.FileAuthManager{
			Region:   rawInput.Login,
			Login:    rawInput.AuthFileProfile,
			Password: rawInput.Password,
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
	err := uploader.PostObject(input.FilePath, input.FileName)
	if err != nil {
		return "", err
	}

	return uploader.GetObjectUrl(input.BucketName, input.ObjectLifeTimeInHours)
}
