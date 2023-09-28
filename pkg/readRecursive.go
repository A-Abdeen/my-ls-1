package Myls

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadRecursive(dir string, flags Flags) error {
	if !flags.R {
		filesAndFolders, err := Read(dir, flags)
		if err != nil {
			return err
		}
		filesAndFolders = sortFilesAndFolders(filesAndFolders, flags)
		for _, file := range filesAndFolders {
			printFileOrDir(file, file.Info.IsDir(), flags)
		}
		fmt.Println()
		return nil
	}

	return filepath.WalkDir(dir, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error at path:", path, "-", err) // Debug print
			return err
		}

		// Skipping if it's a hidden file/directory and -a is not enabled, except for the root directory
		isHidden := entry.Name()[0] == '.'
		if !isRoot && isHidden && !flags.A {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		isRoot = false // Reset the flag after processing the root directory

		// Print directory contents if it's a directory
		if entry.IsDir() {
			fmt.Println(path + ":" + Reset)
			filesAndFolders, err := Read(path, flags)
			if err != nil {
				return err
			}
			filesAndFolders = sortFilesAndFolders(filesAndFolders, flags)
			for _, file := range filesAndFolders {
				printFileOrDir(file, file.Info.IsDir(), flags)
			}
			fmt.Println()
			fmt.Println()
		}
		return nil
	})
}
