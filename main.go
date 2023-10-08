package main

import (
	"fmt"
"strings"
	Myls "Myls/pkg"
)

func main() {
	flags, args := Myls.Parse()
	if len(args) == 0 {
		if !flags.R {
			Myls.NonRecursive(".", flags)
		} else {
			Myls.ReadRecursive(".", flags)
		}
	} else { // checks wether the totally not smart user has entered a directory or a file along with run command(took me 2 hours to figure this out, but 5 minutes to fix it)
		for _, arg := range args {
			if !flags.R {
				Myls.NonRecursive(arg, flags)
				continue
			}
			Myls.ReadRecursive(arg, flags)
		}
	}
	// Myls.Success = Myls.TrimEmptyStrings(Myls.Success)
	for _, fail := range Myls.Fail {
		fmt.Println("myls: cannot access '" + fail + "': No such file or directory")
	}
	if flags.L && !flags.R {
		fmt.Println("total ", Myls.TotalBlocks/2)
	}
	// for i, success := range Myls.Success {
	// 	fmt.Println(success[i], "\t", success[i+1])
	// 	i++
	// }
	for i:=0;i<len(Myls.Success);i++{
		if !flags.L{
		if i >= (len(Myls.Success)/2){
			break
		}
		fmt.Println(Myls.Success[i], strings.Repeat(" ", (55-len(Myls.Success[i]))), Myls.Success[(len(Myls.Success)/2)+i])
	} else {
		fmt.Print(Myls.Success[i])
	}
}	
	// fmt.Println()
}
