package Myls
// Sorts the files and folders by modification time
func sortByModification(filesAndFolders []File, flags Flags) []File {
for i:=0;i<len(filesAndFolders);i++{
	for j:=i+1;j<len(filesAndFolders);j++{
		if flags.Rr{
			if filesAndFolders[i].ModTime.After(filesAndFolders[j].ModTime){
			filesAndFolders[i], filesAndFolders[j] = filesAndFolders[j], filesAndFolders[i]
		} else {
			continue
		}
	}else{
		if filesAndFolders[j].ModTime.After(filesAndFolders[i].ModTime){
			filesAndFolders[i], filesAndFolders[j] = filesAndFolders[j], filesAndFolders[i]
		} else {
			continue
		}
}
	}
}
return filesAndFolders
}