package poc

import (
	"fmt"
	"testing"

	"github.com/expr-lang/expr"
	"github.com/stretchr/testify/require"
)

// Benchmark compares performance of two approaches for expression calculation:
// 1. Using manually constructed Abstract Syntax Trees (AST) and closures (concept approach).
// 2. Compiling expressions using the third-party library github.com/expr-lang/expr.
func Benchmark(b *testing.B) {
	var x, y float64
	x, y = 3, 5

	// Expression to be evaluated using both approaches.
	testExpressionStr := "X+(6*Y)"
	testExpressionParams := struct {
		X float64
		Y *float64
	}{x, &y}

	// Construct the AST corresponding to the test expression.
	conceptAst := Node{
		operator: BinaryOperator[float64, float64, float64](func(x, y float64) float64 {
			return x + y
		}),
		arguments: []any{
			Const[float64]{Value: x}, // Constant value for X
			Node{ // Multiplication subtree
				operator: BinaryOperator[float64, float64, float64](func(x, y float64) float64 {
					return x * y
				}),
				arguments: []any{
					Pointer[float64]{ValuePtr: &y}, // Pointer to variable Y
					Const[float64]{Value: 6},       // Constant multiplier
				},
			},
		},
	}

	// Compile the AST into a callable Go function.
	conceptProg, ok := conceptAst.Build().(func() float64)
	require.True(b, ok)

	// Compile the same expression using the external library github.com/expr-lang/expr.
	exprProg, err := expr.Compile(testExpressionStr, expr.AsFloat64(), expr.Env(testExpressionParams))
	require.Nil(b, err)

	// Wrapper function for executing the expression compiled by the external library.
	exprRun := func() float64 {
		val, _ := expr.Run(exprProg, testExpressionParams)
		return val.(float64)
	}

	// Ensure both methods yield the same result.
	require.Equal(b, conceptProg(), exprRun())

	// Display details about the expression being tested.
	fmt.Printf("expression: %q\n", testExpressionStr)
	fmt.Printf("params:\n\tX=%v\n\t*Y=%v\n", testExpressionParams.X, *testExpressionParams.Y)
	fmt.Printf("concept result: %v\n", conceptProg())
	fmt.Printf("expr result: %v\n\n", exprRun())

	// Benchmark the custom-built AST-based solution.
	b.Run("concept", func(b *testing.B) {
		for b.Loop() {
			conceptProg()
		}
	})

	// Benchmark the external library's compilation approach.
	b.Run("expr", func(b *testing.B) {
		for b.Loop() {
			exprRun()
		}
	})
}
