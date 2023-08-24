package input

import (
	"flag"
	"fmt"
	"github.com/CatInsideBoxUnderTheTable/ITransfer/utils"
	"os"
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

func GetUserConfig() UserInput {
	consoleInput := readInputFromConsole()
	envInput := readInputFromEnvironment()

	return UserInput{
		ObjectLifeTimeInHours: consoleInput.objectLifeTimeInHours,
		FilePath:              consoleInput.filePath,
		FileName:              consoleInput.fileName,
		BucketName:            envInput.bucketName,
		BucketRegion:          envInput.bucketRegion,
		AuthFileProfile:       envInput.authFileProfile,
	}
}

func readInputFromConsole() userConsoleInput {
	validateArgsExisting()
	filePath := os.Args[len(os.Args)-1]
	validateFilePath(filePath)

	defaultFileName := extractFileNameFromPath(filePath)
	fileName := flag.String("n", defaultFileName, "Define own unique file name. Must be unique in storage")
	lifetime := flag.Uint("l", 2, "Define lifespan of link in hours")
	flag.Parse()

	validateStringNotEmpty(*fileName, "provided filename contains white characters")

	return userConsoleInput{
		filePath:              filePath,
		fileName:              *fileName,
		objectLifeTimeInHours: *lifetime,
	}
}

func readInputFromEnvironment() userEnvInput {
	bucketName := readEnvironmentVariableAsString("AWS_STORAGE_BUCKET_NAME")
	bucketRegion := readEnvironmentVariableAsString("AWS_STORAGE_BUCKET_REGION")
	fileAuthProfile := readEnvironmentVariableAsString("AWS_FILE_AUTH_PROFILE")

	return userEnvInput{
		bucketName:      bucketName,
		authFileProfile: fileAuthProfile,
		bucketRegion:    bucketRegion,
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
