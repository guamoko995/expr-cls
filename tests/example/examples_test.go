package example_test

import (
	"fmt"

	exprcls "github.com/guamoko995/expr-cls"

	// Using the example environment
	_ "github.com/guamoko995/expr-cls/tests/example/def_env"
)

// CompileAndCalcExample demonstrates how to compile and evaluate expressions
// using the expr-cls package.
func Example() {
	// Define a data structure containing input variables for our expression.
	type InputData struct {
		A int
		B int
	}

	// Parse and compile the expression "3 + A * B".
	prog, err := exprcls.Compile[InputData, float64]("3 + A * -B + pi")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Provide input data for evaluating the expression.
	input := InputData{A: 7, B: 10}

	// Evaluate the expression using the provided input data.
	result := prog(input)

	// Print the computed result.
	fmt.Println(result)

	// Output: -63.8584073464102
}
