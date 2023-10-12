package Myls

func FindSize(file *File) {
	if len(file.Group) > len(Size.Group) {
		Size.Group = file.Group
	}
	if len(file.Owner) > len(Size.Owner) {
		Size.Owner = file.Owner
	}
	if file.Links > Size.Links {
		Size.Links = file.Links
	}
	if len(file.Permissions.String()) > len(Size.Permissions.String()){
		Size.Permissions = file.Permissions
	}
	if file.Size > Size.Size {
		Size.Size = file.Size
	}
}