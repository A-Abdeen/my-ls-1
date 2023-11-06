package Myls

import (
	"os"
)

func symLinkFunc(file File, flags Flags) string {
	followLink := false
	// originFile := ""
	originFile, err := os.Readlink(file.Info.Name())
	if err != nil {
		originFile, err = os.Readlink("/usr/bin/" + file.Info.Name())
		originFile = Green + originFile + Reset
		if err != nil {
			originFile, _ = os.Readlink("/dev/" + file.Info.Name())
			originFile = Yellow + originFile + Reset
		}
	}
	// originFile2 := os.Symlink(file.Info.Name(), "pkg")
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == file.Info.Name() {
			followLink = true
		}
	}
	for _, file23 := range FilesAndFolders23 {
		if originFile == file23.Info.Name() && file23.Info.IsDir() {
			// fmt.Println(file23)
			if !flags.L && followLink {
				NonRecursive(file23.Info.Name(), flags)
				return ""
			} else {
				originFile = Blue + originFile + Reset
			}
		}
	}
	originFile = " -> " + originFile
	return originFile
}
