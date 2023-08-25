package input

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type UserInput struct {
	ObjectLifeTimeInHours uint
	FilePath              string
	FileName              string
	BucketName            string
	BucketRegion          string
	AuthFileProfile       string
}

type userConsoleInput struct {
	objectLifeTimeInHours uint
	filePath              string
	fileName              string
}

type userEnvInput struct {
	bucketName      string
	bucketRegion    string
	authFileProfile string
}

func GetUserConfig() (UserInput, error) {
	consoleInput, err := readInputFromConsole()
	if err != nil {
		return UserInput{}, err
	}

	envInput, err := readInputFromEnvironment()
	if err != nil {
		return UserInput{}, err
	}

	return UserInput{
		ObjectLifeTimeInHours: consoleInput.objectLifeTimeInHours,
		FilePath:              consoleInput.filePath,
		FileName:              consoleInput.fileName,
		BucketName:            envInput.bucketName,
		BucketRegion:          envInput.bucketRegion,
		AuthFileProfile:       envInput.authFileProfile,
	}, nil
}

func readInputFromConsole() (userConsoleInput, error) {
	if !argsExisting() {
		return userConsoleInput{}, errors.New("provided console arguments are invalid")
	}

	absoluteFilePath, err := filepath.Abs(os.Args[len(os.Args)-1])
	if err != nil {
		return userConsoleInput{}, err
	}

	if err != nil {
		return userConsoleInput{}, err
	}

	fileName := flag.String("n", filepath.Base(absoluteFilePath), "Define file name")
	lifetime := flag.Uint("l", 2, "Define lifespan of link in hours")
	flag.Parse()

	return userConsoleInput{
		filePath:              absoluteFilePath,
		fileName:              *fileName,
		objectLifeTimeInHours: *lifetime,
	}, nil
}

func readInputFromEnvironment() (userEnvInput, error) {
	bucketName, err := readEnvironmentVariableAsString("AWS_STORAGE_BUCKET_NAME")
	if err != nil {
		return userEnvInput{}, err
	}

	bucketRegion, err := readEnvironmentVariableAsString("AWS_STORAGE_BUCKET_REGION")
	if err != nil {
		return userEnvInput{}, err
	}

	fileAuthProfile, err := readEnvironmentVariableAsString("AWS_FILE_AUTH_PROFILE")
	if err != nil {
		return userEnvInput{}, err
	}

	return userEnvInput{
		bucketName:      bucketName,
		authFileProfile: fileAuthProfile,
		bucketRegion:    bucketRegion,
	}, nil
}

func readEnvironmentVariableAsString(envVarName string) (string, error) {
	val, exists := os.LookupEnv(envVarName)

	if !exists || stringEmpty(val) {
		return "", fmt.Errorf("required environment variable %s is invalid", envVarName)
	}

	return val, nil
}

func argsExisting() bool {
	return len(os.Args) >= 1
}

func stringEmpty(val string) bool {
	if len(val) == 0 {
		return true
	}

	for _, char := range val {
		if char != ' ' {
			return false
		}
	}

	return true
}
