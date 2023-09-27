// main.go
package main

import (
	Myls "Myls/pkg"
	"fmt"
)

func main() {
	flags := Myls.Parse()
	files, err := Myls.Read(".", flags.A)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Printf("Flags: %+v\n", flags)
	fmt.Println("Files:")
	for _, file := range files {
		fmt.Println(file.Info.Name())
	}
}
