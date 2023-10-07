package Myls

// Sorts the files and folders by alphabetical order
func sortFilesAndFolders(filesAndFolders []File, flags Flags) []File {
	for i:=0;i<len(filesAndFolders);i++{
		for j:=i+1;j<len(filesAndFolders);j++{
			for k:=0;k<len(filesAndFolders[i].Info.Name());k++{
			if !flags.Rr{
				if filesAndFolders[i].Info.Name()[k] > filesAndFolders[j].Info.Name()[k]{
				filesAndFolders[i], filesAndFolders[j] = filesAndFolders[j], filesAndFolders[i]
				break
			}else if filesAndFolders[i].Info.Name()[k] == filesAndFolders[j].Info.Name()[k] {
				continue
			} else {
				break
			}
		}else{
			if filesAndFolders[i].Info.Name()[k] < filesAndFolders[j].Info.Name()[k]{
				filesAndFolders[i], filesAndFolders[j] = filesAndFolders[j], filesAndFolders[i]
				break
			}else if filesAndFolders[i].Info.Name()[k] == filesAndFolders[j].Info.Name()[k] {
				continue
			} else {
				break
			}
		}}
		}
	}
	return filesAndFolders
}
