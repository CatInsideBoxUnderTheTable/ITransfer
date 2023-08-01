package utils

import (
	"fmt"
	"os"
)

func ReadFile(path string) *os.File {
	file, err := os.Open(path)
	PanicIfErr(err, fmt.Sprintf("file at path %s cannot be opened. INNER ERR: %s", path, err))

	return file
}

func CloseFile(file *os.File) {
	err := file.Close()
	PanicIfErr(err, fmt.Sprintf("file at path %s cannot be closed. INNER ERR: %s", file.Name(), err))
}
