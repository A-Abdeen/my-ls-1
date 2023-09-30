package Myls

import (
	"fmt"
	"os"
)

// Read the directory and returns a slice of File structs
func Read(dir string, flag Flags) ([]File, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filesAndFolders []File
	for _, entry := range entries {
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
