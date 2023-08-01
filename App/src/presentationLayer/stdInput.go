package presentation

import (
	"flag"
	"fmt"
	presentationModels "github.com/aws/aws-sdk-go/service/s3/App/src/presentationLayer/models"
	"github.com/aws/aws-sdk-go/service/s3/App/src/utils"
	"os"
)

func ReadInputFromConsole() presentationModels.UserInput {
	validateArgsExisting()
	filePath := os.Args[len(os.Args)-1]
	validateFilePath(filePath)

	defaultFileName := extractFileNameFromPath(filePath)
	fileName := flag.String("n", defaultFileName, "Define own unique file name. Must be unique in storage")
	lifetime := flag.Uint("l", 2, "Define lifespan of link in hours")
	flag.Parse()

	validateFileName(*fileName)

	return presentationModels.UserInput{
		FilePath:              &filePath,
		FileName:              fileName,
		ObjectLifeTimeInHours: lifetime,
	}
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

func validateFileName(fileName string) {
	for _, char := range fileName {
		if char == ' ' {
			panic("provided filename contains white characters")
		}
	}
}
