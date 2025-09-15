package poc

import (
	"fmt"
	"testing"

	"github.com/expr-lang/expr"
	exprcls "github.com/guamoko995/expr-cls"
	"github.com/stretchr/testify/require"
)

// Benchmark compares performance of two approaches for expression calculation:
// 1. Using manually constructed Abstract Syntax Trees (AST) and closures (concept approach).
// 2. Compiling expressions using the third-party library github.com/expr-lang/expr.
func Benchmark(b *testing.B) {
	type input struct {
		X float64
		Y float64
	}

	exprcls.RegisterSource[input]()

	// Expression to be evaluated using both approaches.
	testExpressionStr := "X+(6*Y)"

	// Compile the AST into a callable Go function.
	conceptProg, err := exprcls.Compile[input, float64](testExpressionStr)
	require.Nil(b, err)

	// Compile the same expression using the external library github.com/expr-lang/expr.
	exprProg, err := expr.Compile(testExpressionStr, expr.AsFloat64(), expr.Env(input{}))
	require.Nil(b, err)

	// Wrapper function for executing the expression compiled by the external library.
	exprRun := func(in input) float64 {
		val, _ := expr.Run(exprProg, in)
		return val.(float64)
	}

	testExpressionParams := input{X: 3, Y: 5}
	// Ensure both methods yield the same result.
	require.Equal(b, conceptProg(testExpressionParams), exprRun(testExpressionParams))

	// Display details about the expression being tested.
	fmt.Printf("expression: %q\n", testExpressionStr)
	fmt.Printf("params:\n\tX=%v\n\t*Y=%v\n", testExpressionParams.X, testExpressionParams.Y)
	fmt.Printf("concept result: %v\n", conceptProg(testExpressionParams))
	fmt.Printf("expr result: %v\n\n", exprRun(testExpressionParams))

	// Benchmark the custom-built AST-based solution.
	b.Run("concept", func(b *testing.B) {
		for b.Loop() {
			conceptProg(testExpressionParams)
		}
	})

	// Benchmark the external library's compilation approach.
	b.Run("expr", func(b *testing.B) {
		for b.Loop() {
			exprRun(testExpressionParams)
		}
	})
}
