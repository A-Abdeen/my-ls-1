package Myls

import (
	"fmt"
	"os"
)

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
			fmt.Println("Error:", err) // Adjust as needed
			continue
		}

		filesAndFolders = append(filesAndFolders, file)
	}

	return filesAndFolders, nil
}
