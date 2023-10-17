package Myls

import "strings"
func FindSize(file *File) {
	stringPermissions := file.Permissions.String()
	stringPermissions = strings.TrimPrefix(stringPermissions, "D")
	if len(file.Group) > len(Size.Group) {
		Size.Group = file.Group
	}
	if len(file.Owner) > len(Size.Owner) {
		Size.Owner = file.Owner
	}
	if file.Links > Size.Links {
		Size.Links = file.Links
	}
	if len(file.Permissions.String()) > len(stringPermissions) {
		Size.Permissions = file.Permissions
	}
	if file.Size > Size.Size {
		Size.Size = file.Size
	}
}
