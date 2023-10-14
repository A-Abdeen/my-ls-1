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
	OriginFile, _ := os.Readlink(dir)
	if OriginFile != "" {
		entries, _ = os.ReadDir(".")
			file = dir
	}
	if err != nil {
		if strings.ContainsRune(dir, '/'){			
		dir, file, _  = strings.Cut(dir, "/")
		entries, _ = os.ReadDir(dir)
	} else {
			entries, _ = os.ReadDir(".")
			file = dir
	}
}
	for _, entry := range entries {
		// fmt.Println(entry)
		// fmt.Println(file)
		if file != "" && (file) != entry.Name() {
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
		FilesAndFolders = append(FilesAndFolders, file)
	}
	if FilesAndFolders == nil {
		return nil, err
	}
	return FilesAndFolders, nil
	
}