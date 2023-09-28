package main

import (
	Myls "Myls/pkg"
)

func main() {
	flags := Myls.Parse()
	Myls.ReadRecursive(".", flags) // Passing flags to ReadRecursive
}
