package presentation

import (
	"flag"
	"fmt"
	"os"

	presentationModels "github.com/aws/aws-sdk-go/service/s3/App/src/presentationLayer/models"
	"github.com/aws/aws-sdk-go/service/s3/App/src/utils"
)

func GetUserConfig() presentationModels.UserInput {
	consoleInput := readInputFromConsole()
	envInput := readInputFromEnvironment()

	return presentationModels.UserInput{
		UserEnvInput:     &envInput,
		UserConsoleInput: &consoleInput,
	}
}

func readInputFromConsole() presentationModels.UserConsoleInput {
	validateArgsExisting()
	filePath := os.Args[len(os.Args)-1]
	validateFilePath(filePath)

	defaultFileName := extractFileNameFromPath(filePath)
	fileName := flag.String("n", defaultFileName, "Define own unique file name. Must be unique in storage")
	lifetime := flag.Uint("l", 2, "Define lifespan of link in hours")
	flag.Parse()

	validateStringNotEmpty(*fileName, "provided filename contains white characters")

	return presentationModels.UserConsoleInput{
		FilePath:              &filePath,
		FileName:              fileName,
		ObjectLifeTimeInHours: lifetime,
	}
}

func readInputFromEnvironment() presentationModels.UserEnvInput {
	bucketName := readEnvironmentVariableAsString("AWS_STORAGE_BUCKET_NAME")
	bucketRegion := readEnvironmentVariableAsString("AWS_STORAGE_BUCKET_REGION")
	fileAuthProfile := readEnvironmentVariableAsString("AWS_FILE_AUTH_PROFILE")

	return presentationModels.UserEnvInput{
		BucketName:      &bucketName,
		AuthFileProfile: &fileAuthProfile,
		BucketRegion:    &bucketRegion,
	}
}
func readEnvironmentVariableAsString(envVarName string) string {
	val, exists := os.LookupEnv(envVarName)

	if !exists {
		panic(fmt.Sprintf("environment variable '%s' not set", envVarName))
	}
	validateStringNotEmpty(val, fmt.Sprintf("environment variable '%s' is empty", envVarName))

	return val
}

func extractFileNameFromPath(filePath string) string {
	info, err := os.Stat(filePath)
	utils.PanicIfErr(err, fmt.Sprintf("Unable to open file. INNER ERR: %s", err))

	return info.Name()
}

func validateArgsExisting() {
	if len(os.Args) <= 1 {
		panic(fmt.Sprintf("Provided invalid number of arguments. No: %d", len(os.Args)))
	}
}

func validateFilePath(filePath string) {
	info, err := os.Stat(filePath)
	utils.PanicIfErr(err, fmt.Sprintf("Unable to open file. INNER ERR: %s", err))

	if info.IsDir() {
		panic("Unable to send directory. Please, compress it first")
	}
}

func validateStringNotEmpty(fileName string, errMessage string) {
	for _, char := range fileName {
		if char == ' ' {
			panic(errMessage)
		}
	}
}
