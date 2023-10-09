package Myls

import (
	"fmt"
)

func ReadRecursive(rootDir string, flags Flags) error {
	// Helper function to print entries in a directory
	
	printEntries := func(dir string, entries []File) {
		for _, entry := range entries {
			printFileOrDir(entry, entry.Info.IsDir(), flags)
		}
		// fmt.Println()
		// fmt.Println()
		Success = append(Success, "\n\n")
	}
	// Recursive function to process directories
	var processDir func(dir string) error
	processDir = func(dir string) error {
		entries, err := Read(dir, flags)
		if err != nil {
			Fail = append(Fail, dir)
			// fmt.Print(err)
			return err
		}

		// Print the directory name
		// fmt.Println(dir + ":")
		Success = append(Success, dir+":\n")

		// Print the directory entries
		entries = sortFilesAndFolders(entries, flags)
		if flags.T {
			entries = sortByModification(entries, flags)
		}
		var blocks int64
		if flags.L {
			for _, file := range entries {
				blocks += file.blockSize
			}
			Success = append(Success, "total "+fmt.Sprint(blocks/2)+"\n")
		}
		
		printEntries(dir, entries)

		// Find subdirectories and process them recursively
		for _, entry := range entries {
			if entry.Info.IsDir() {
				subDir := dir + "/"+entry.Info.Name()
				if err := processDir(subDir); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return processDir(rootDir)
}
