package Myls

import (
	"fmt"
	"io/fs"
	"os"
)

func FlagA(dir string) []fs.DirEntry {
	var filenumber int
	var filenumber2 int
	var newfileDirFinal []fs.DirEntry
	path, err := os.Getwd()
	if dir != "" && dir != "." {
		path = path + "/" + dir
	}
	path, filename := MySplit(path, "/")
	path2, filename2 := MySplit(path, "/")
	if err != nil {
		fmt.Println(err)
	}
	newfileDir, err := os.ReadDir(path)
	for i := 0; i < len(newfileDir); i++ {
		if newfileDir[i].Name() == filename {
			filenumber = i
			break
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	newfileDir2, err2 := os.ReadDir(path2)
	for i := 0; i < len(newfileDir2); i++ {
		if newfileDir2[i].Name() == filename2 {
			filenumber2 = i
			break
		}
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	newfileDirFinal = append(newfileDirFinal, newfileDir[filenumber])
	newfileDirFinal = append(newfileDirFinal, newfileDir2[filenumber2])
	return newfileDirFinal
}
