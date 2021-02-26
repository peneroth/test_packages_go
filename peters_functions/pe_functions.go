package petersfunctions

import "fmt"

// Vertex contains X and Y coordinates
type Vertex struct {
	X int
	Y int
}

// TestFunction returns nothing
func TestFunction() {
	// Comment 2
	fmt.Println("Print from package petersfunctions, TestFunction()")
	internalTestFunction()
}

func internalTestFunction() {
	fmt.Println("Print from package petersfunctions, internalTestFunction()")
}

// AddOneToVertex adds 1 to X and Y value, no pointer
func AddOneToVertex(v Vertex) {
	v.X = v.X + 1
	v.Y = v.Y + 1
}

// AddOneToVertex adds 1 to X and Y value, use pointer
func AddOneToVertexp(v *Vertex) {
	v.X = v.X + 1
	v.Y = v.Y + 1
}
