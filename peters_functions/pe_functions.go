package petersfunctions

import "fmt"

// TestFunction returns nothing
func TestFunction() {
	// Comment 2
	fmt.Println("Print from package petersfunctions, TestFunction()")
	internalTestFunction()
}

func internalTestFunction() {
	fmt.Println("Print from package petersfunctions, internalTestFunction()")
}
