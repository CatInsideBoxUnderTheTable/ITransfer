package dataLayerAbstractions

type BucketUploader interface {
	InitializeSession(bucketName string)
	PostObject(filePath string, objectName string)
	GetObjectUrl(objectName string, urlExpirationHours uint) string
	Close()
}
