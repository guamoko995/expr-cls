package expr

import (
	"fmt"
)

// ExampleExpr demonstrates how to build and evaluate expressions using the expr
// package.
func Example() {
	// Define a data structure containing input variables for our expression.
	type InputData struct {
		A int
		B int
	}

	// Register the data structure in the execution environment.
	RegisterSource[InputData]()

	// Parse and compile the expression "3 + A * B".
	fn, err := Compile[InputData, int]("3 + A * -B")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Provide input data for evaluating the expression.
	input := InputData{A: 7, B: 10}

	// Evaluate the expression using the provided input data.
	result := fn(input)

	// Print the computed result.
	fmt.Println(result)

	// Output: -67
}
