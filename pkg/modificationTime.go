package Myls

import "sort"

// Sorts the files and folders by modification time
func sortByModification(filesAndFolders []File, flags Flags) []File {
	sort.Slice(filesAndFolders, func(i, j int) bool {
		// Check if the -t flag is set
		return filesAndFolders[i].ModTime.After(filesAndFolders[j].ModTime)
	})
	return filesAndFolders
}
