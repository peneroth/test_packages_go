package main

import (
	"fmt"

	petersfunctions "github.com/peneroth/test_packages_go/peters_functions"
)

// Test globalVariable, accessable within package main
var globalVariable int = 1

func main() {
	fmt.Println("Main package - main file")

	// Note, capital first letter not needed when in same package but in another file
	aMainFunction()
	petersfunctions.TestFunction()

	// Test globalVariable
	globalVariable = 3
	// Global variables can be reached within the same package, but is perhaps not very nice
	printGlobalVar()

	// Test pointer
	a := 1
	p := &a
	incOneViaPointer(p)
	fmt.Println("a = ", *p)

	// Test struct in package
	v := petersfunctions.Vertex{X: 1, Y: 2}
	// Call with structure => changes are only available in local function
	petersfunctions.AddOneToVertex(v)
	fmt.Println("From main function: X = ", v.X, "Y = ", v.Y)
	// Call with pointer to structure => changes are maintained in caller function
	petersfunctions.AddOneToVertexp(&v)
	fmt.Println("From main function: X = ", v.X, "Y = ", v.Y)

}
