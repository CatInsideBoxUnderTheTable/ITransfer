package storage

type BucketUploader interface {
	InitializeSession(bucketName string)
	PostObject(filePath string, objectName string)error
	GetObjectUrl(objectName string, urlExpirationHours uint) (string, error)
	Close()
}

type AuthManager interface {
	OpenSession()
	GetExistingSession() any
	Close()
}
