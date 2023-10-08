package Myls

import (
	"fmt"
	"os"
	"strings"
)

// Read the directory and returns a slice of File structs
func Read(dir string, flag Flags) ([]File, error) {
	var file string
	entries, err := os.ReadDir(dir)
	if err != nil {
		if strings.ContainsRune(dir, '/'){
		dir, file, _  = strings.Cut(dir, "/")
	} else {
			entries, _ = os.ReadDir(".")
			file = dir
	}
}
	var filesAndFolders []File
	for _, entry := range entries {
		if file != "" && file != entry.Name(){
			continue
		}
		if !flag.A && entry.Name()[0] == '.' {
			continue
		}

		file := File{Info: entry}
		err := file.PopulateInfo()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		filesAndFolders = append(filesAndFolders, file)
	}

	return filesAndFolders, nil

}