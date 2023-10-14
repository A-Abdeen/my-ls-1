package Myls

import (
	"os"
	"strings"
)
func FFlagFunc(file File) string {
	fileName := file.Info.Name()
	if file.Info.IsDir() {
		fileName = fileName + "/"
	} else if strings.ContainsRune(fileName, '.'){
		fileName = fileName + "*"
	} else if originFile, _ := os.Readlink(fileName); originFile != "" {
		fileName = fileName + "@"
	}
	return fileName
}
