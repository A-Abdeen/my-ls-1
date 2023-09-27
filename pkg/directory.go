package Myls

import (
	"os"
)

// Read function to read the directory and return files
func Read(dir string, showHidden bool) ([]File, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []File
	for _, entry := range entries {
		if !showHidden && entry.Name()[0] == '.' {
			continue
		}

		var file File
		file.Info = entry
		files = append(files, file)
	}
	return files, nil
}
