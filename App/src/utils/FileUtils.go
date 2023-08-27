package utils

import "os"

func ReadFile(path string) (*os.File, error) {
	return os.Open(path)
}

func CloseFile(file *os.File) {
	file.Close()
}
