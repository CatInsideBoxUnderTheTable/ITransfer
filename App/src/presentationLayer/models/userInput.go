package presentationModels

type UserInput struct {
	UserConsoleInput *UserConsoleInput
	UserEnvInput     *UserEnvInput
}
type UserConsoleInput struct {
	ObjectLifeTimeInHours *uint
	FilePath              *string
	FileName              *string
}

type UserEnvInput struct {
	BucketName      *string
	BucketRegion    *string
	AuthFileProfile *string
}
