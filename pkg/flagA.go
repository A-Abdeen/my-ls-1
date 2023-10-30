package Myls

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func FlagA() ([]fs.DirEntry) {
	var filenumber int
	var filenumber2 int
	var newfileDirFinal []fs.DirEntry
	path, err := os.Getwd()
	patharray := strings.Split(path, "/")
	fmt.Println(patharray)
	if err != nil {
		fmt.Println(err)
	}
	if len(patharray) >2 {
	newfileDir, err := os.ReadDir("/" + patharray[(len(patharray)-3)] + "/" + patharray[(len(patharray)-2)])
	for i := 0; i < len(newfileDir); i++ {
		filename := newfileDir[i].Name()
		if patharray[(len(patharray)-1)] == filename {
			filenumber = i
			break
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	newfileDir2, err2 := os.ReadDir("/" + patharray[(len(patharray)-3)])
	for i := 0; i < len(newfileDir2); i++ {
		filename := newfileDir2[i].Name()
		if patharray[(len(patharray)-2)] == filename {
			filenumber2 = i
			break
		}
	}
	if err2 != nil {
		fmt.Println(err)
	}
	newfileDirFinal = append(newfileDirFinal, newfileDir[filenumber])
	newfileDirFinal = append(newfileDirFinal, newfileDir2[filenumber2])}
	return newfileDirFinal
}
