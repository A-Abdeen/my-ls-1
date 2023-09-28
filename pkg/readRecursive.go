package Myls

import (
	"fmt"
	"path/filepath"
)

func ReadRecursive(rootDir string, flags Flags) error {
	// Helper function to print entries in a directory
	printEntries := func(dir string, entries []File) {
		for _, entry := range entries {
			printFileOrDir(entry, entry.Info.IsDir(), flags)
		}
		fmt.Println()
		fmt.Println()
	}

	// Recursive function to process directories
	var processDir func(dir string) error
	processDir = func(dir string) error {
		entries, err := Read(dir, flags)
		if err != nil {
			Fail = append(Fail, dir)
			return err
		}

		// Print the directory name
		fmt.Println(dir + ":")

		// Print the directory entries
		entries = sortFilesAndFolders(entries, flags)
		printEntries(dir, entries)

		// Find subdirectories and process them recursively
		for _, entry := range entries {
			if entry.Info.IsDir() {
				subDir := filepath.Join(dir, entry.Info.Name())
				if err := processDir(subDir); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return processDir(rootDir)
}
