package main

import (
	"fmt"

	petersfunctions "github.com/peneroth/test_packages_go/peters_functions"
)

func main() {
	fmt.Println("Main package - main file")

	aMainFunction()

	petersfunctions.TestFunction()
}
