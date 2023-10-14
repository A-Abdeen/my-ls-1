package Myls

import (
	"strings"
)

// Sorts the files and folders by alphabetical order
func sortFilesAndFolders(filesAndFolders []File, flags Flags) []File {
	for i := 0; i < len(filesAndFolders); i++ {
		for j := i + 1; j < len(filesAndFolders); j++ {
				x := (filesAndFolders[j].Info.Name())
				y := (filesAndFolders[i].Info.Name())
				x = strings.ToLower(x)
				y = strings.ToLower(y)
				x = Alphanumeric(x)
				y = Alphanumeric(y)
				for k := 0; k < len(x) && k < len(y); k++ {
				if !flags.Rr {
					if y[k] > x[k] {
						filesAndFolders[i], filesAndFolders[j] = filesAndFolders[j], filesAndFolders[i]
						break
					} else if x[k] == y[k] {
						continue
					} else {
						break
					}
				} else {
					if y[k] < x[k] {
						filesAndFolders[i], filesAndFolders[j] = filesAndFolders[j], filesAndFolders[i]
						break
					} else if x[k] == y[k] {
						continue
					} else {
						break
					}
				}
			}
		}
	}
	return filesAndFolders
}
