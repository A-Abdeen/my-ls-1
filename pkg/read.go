package Myls
import (
	"fmt"
	"os"
	"strings"
)
// Read the directory and returns a slice of File structs
func Read(dir string, flag Flags) ([]File, error) {
	var singlefile string
	entries, err := os.ReadDir(dir)
	OriginFile, _ := os.Readlink(dir)
	if OriginFile != "" {
		entries, _ = os.ReadDir(".")
			singlefile = dir
	}
	if err != nil {
		if strings.ContainsRune(dir, '/'){			
		dir, singlefile, _  = strings.Cut(dir, "/")
		entries, _ = os.ReadDir(dir)
	} else if strings.ContainsRune(dir, '.') && FilesAndFolders == nil{
			entries, _ = os.ReadDir(".")
			singlefile = dir
	}
}
var filesAndFolders []File
	for _, entry := range entries {
		// fmt.Println(entry)
		// fmt.Println(file)
		file := File{Info: entry}
		err := file.PopulateInfo()
		FilesAndFolders23 = append(FilesAndFolders23, file)
		if singlefile != "" && (singlefile) != entry.Name() {
			continue
		}
		if !flag.A && entry.Name()[0] == '.' {
			continue
		}
		
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		filesAndFolders = append(filesAndFolders, file)
	}
	if filesAndFolders == nil {
		return nil, err
	}
	FilesAndFolders = filesAndFolders
	return FilesAndFolders, nil
	
}