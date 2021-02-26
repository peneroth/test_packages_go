package main

// Functions in package main can be placed in several files, belonging to the same package

import "fmt"

func aMainFunction() {
	// Comment 2
	fmt.Println("Print from package main, aMainFunction()")
}

func printGlobalVar() {
	fmt.Println("globalVariable = ", globalVariable)
}

func incOneViaPointer(i *int) {
	*i = *i + 1
}
