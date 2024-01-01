package tango_errors

import "fmt"

func Debug(name, value string) {

	fmt.Println("")
	fmt.Println("--------------------------------")
	fmt.Println(" DEBUG: ")
	fmt.Printf("%s: %s \n", name, value)
	fmt.Println("--------------------------------")
	fmt.Println("")

}
