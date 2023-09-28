package Myls

import (
	"sort"
)

func sortFilesAndFolders(filesAndFolders []File, flags Flags) []File {
	sort.Slice(filesAndFolders, func(i, j int) bool {
		if flags.Rr {
			return filesAndFolders[i].Info.Name() > filesAndFolders[j].Info.Name()
		}
		return filesAndFolders[i].Info.Name() < filesAndFolders[j].Info.Name()
	})
	return filesAndFolders
}
