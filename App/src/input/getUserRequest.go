package input

import (
    "errors"
    "flag"
    "fmt"
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

func GetUserConfig() (UserInput, error){
    consoleInput, err := readInputFromConsole()
    if err != nil{
        return UserInput{}, err
    }

    envInput, err := readInputFromEnvironment()
    if err != nil{
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

func readInputFromConsole() (userConsoleInput,error) {
    if !argsExisting(){
        return userConsoleInput{}, errors.New("proved console arguments are invalid")
    }

    filePath := os.Args[len(os.Args)-1]
    if !validFilePath(filePath){
        return userConsoleInput{}, errors.New("provided file path is invalid")
    }

    defaultFileName, err := extractFileNameFromPath(filePath)
    if err !=nil {
        return userConsoleInput{}, err 
    }

    fileName := flag.String("n", defaultFileName, "Define file name")
    lifetime := flag.Uint("l", 2, "Define lifespan of link in hours")
    flag.Parse()

    if stringEmpty(*fileName){
        return userConsoleInput{}, errors.New("provided file name is invald") 
    }

    return userConsoleInput{
        filePath:              filePath,
        fileName:              *fileName,
        objectLifeTimeInHours: *lifetime,
    }, nil 
}

func readInputFromEnvironment() (userEnvInput, error) {
    bucketName, err:= readEnvironmentVariableAsString("AWS_STORAGE_BUCKET_NAME")
    if err!= nil{
        return userEnvInput{}, err
    }

    bucketRegion, err:= readEnvironmentVariableAsString("AWS_STORAGE_BUCKET_REGION")
    if err!= nil{
        return userEnvInput{}, err
    }

    fileAuthProfile, err:= readEnvironmentVariableAsString("AWS_FILE_AUTH_PROFILE")
    if err!= nil{
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

    if !exists || stringEmpty(val){
        return "", fmt.Errorf("required environment variable %s is invalid", envVarName)    
    }

    return val, nil
}

func extractFileNameFromPath(filePath string) (string, error) {
    info, err := os.Stat(filePath)

    return info.Name(), err
}

func argsExisting() bool{
    return len(os.Args) <= 1
}

func validFilePath(filePath string) bool{
    info, err := os.Stat(filePath)

    if info.IsDir() {
        return false
    }

    return err == nil
}

func stringEmpty(val string) bool {
    var isEmpty bool 

    if len(val) == 0{
        isEmpty = true 
    }

    for _, char := range val {
        if char != ' ' {
            isEmpty = true
        }
    }

    return isEmpty
}
