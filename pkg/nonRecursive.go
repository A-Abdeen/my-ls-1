package Myls
// import "fmt"
func NonRecursive(dir string, flags Flags) {
	// Read the directory
	FilesAndFolders, err := Read(dir, flags)
	if err != nil || dir == "-" {
		Fail = append(Fail, dir)
		return
	}
	// Sort the files and folders
	sortFilesAndFolders(FilesAndFolders, flags)
	// Sort by modification time if -t flag is set
	if flags.T {
		sortByModification(FilesAndFolders, flags)
	}
	// Print the files and folders
	for _, file := range FilesAndFolders {
		printFileOrDir(file, flags)
	}
	// fmt.Println()
	for _, file := range FilesAndFolders {
		TotalBlocks += file.blockSize
	}
}