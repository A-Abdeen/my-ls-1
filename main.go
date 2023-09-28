package main

import (
	"fmt"

	Myls "Myls/pkg"
)

func main() {
	flags, args := Myls.Parse()
	if len(args) == 0 {
		if !flags.R {
			Myls.NonRecursive(".", flags)
			return
		}
		Myls.ReadRecursive(".", flags) // Passing flags to ReadRecursive
		return
	}
	for _, arg := range args {
		if !flags.R {
			Myls.NonRecursive(arg, flags)
			continue
		}
		Myls.ReadRecursive(arg, flags) // Passing flags to ReadRecursive
	}
	for _, fail := range Myls.Fail {
		fmt.Println("myls: cannot access '" + fail + "': No such file or directory")
	}
}
