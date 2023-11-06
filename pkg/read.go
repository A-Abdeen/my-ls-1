package Myls

import (
	"fmt"
	"os"
	"strings"
)

// Read the directory and returns a slice of File structs
func Read(dir string, flag Flags) ([]File, error) {
	var err2 error
	var singlefile string
	entries, err := os.ReadDir(dir)
	OriginFile, _ := os.Readlink(dir)
	if OriginFile != "" {
		entries, _ = os.ReadDir(".")
		singlefile = dir
	} else {
		if err != nil {
			if strings.ContainsRune(dir, '/') {
				dir, singlefile, _ = strings.Cut(dir, "/")
				entries, err2 = os.ReadDir(dir)
				if err2 != nil {
					fmt.Print("myls: cannot access '" + dir + "/': Not a directory\n")
					return nil, nil
				}
			} else if strings.ContainsRune(dir, '.') && FilesAndFolders == nil {
				entries, _ = os.ReadDir(".")
				singlefile = dir
			} else {
				fmt.Print("myls: cannot access '" + dir + "/': No such file or directory\n")
				return nil, err
			}
		}
	}
	if flag.A {
		entries = append(entries, FlagA(dir)...)
	}
	var filesAndFolders []File

	for i, entry := range entries {
		file := File{Info: entry}
		if flag.A && i == len(entries)-1 {
			file.Name = ".."
		} else if flag.A && i == len(entries)-2 {
			file.Name = "."
		} else {
			file.Name = file.Info.Name()
		}
		err := file.PopulateInfo()
		FilesAndFolders23 = append(FilesAndFolders23, file)
		if OriginFile != "" {
			singlefile = dir
		}
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
