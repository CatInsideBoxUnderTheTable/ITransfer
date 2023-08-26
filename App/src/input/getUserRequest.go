package input

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/term"
	"os"
	"path/filepath"
)

type UserInput struct {
	ObjectLifeTimeInHours uint
	FilePath              string
	FileName              string
	BucketName            string
	Login                 string
	Password              string
	AuthFileProfile       string
}

type userConsoleInput struct {
	objectLifeTimeInHours uint
	filePath              string
	fileName              string
	password              string
}

type userEnvInput struct {
	bucketName   string
	bucketRegion string
	login        string
}

func GetUserConfig() (UserInput, error) {
	envInput, err := readInputFromEnvironment()
	if err != nil {
		return UserInput{}, err
	}

	consoleInput, err := readInputFromConsole()
	if err != nil {
		return UserInput{}, err
	}

	return UserInput{
		ObjectLifeTimeInHours: consoleInput.objectLifeTimeInHours,
		FilePath:              consoleInput.filePath,
		FileName:              consoleInput.fileName,
		Password:              consoleInput.password,
		BucketName:            envInput.bucketName,
		Login:                 envInput.bucketRegion,
		AuthFileProfile:       envInput.login,
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

	fileName := flag.String("n", filepath.Base(absoluteFilePath), "Define file name")
	lifetime := flag.Uint("l", 2, "Define lifespan of link in hours")
	flag.Parse()

	password, err := readPassword()
	if err != nil {
		return userConsoleInput{}, err
	}

	return userConsoleInput{
		filePath:              absoluteFilePath,
		fileName:              *fileName,
		objectLifeTimeInHours: *lifetime,
		password:              password,
	}, nil
}

func readPassword() (string, error) {
	fmt.Print("provide console password: ")
	rawPass, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	convertedPass := string(rawPass)
	return convertedPass, nil
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

	login, err := readEnvironmentVariableAsString("AWS_CONSOLE_LOGIN")
	if err != nil {
		return userEnvInput{}, err
	}

	return userEnvInput{
		bucketName:   bucketName,
		login:        login,
		bucketRegion: bucketRegion,
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
