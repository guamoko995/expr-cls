package poc

import (
	"fmt"
	"testing"

	"github.com/expr-lang/expr"
	exprcls "github.com/guamoko995/expr-cls"
	"github.com/stretchr/testify/require"
)

// Benchmark compares performance of two approaches for expression calculation:
// 1. Using expr-cls: full-featured expression compilation from string to Go function (MVP build).
// 2. Compiling expressions using the third-party library github.com/expr-lang/expr.
func Benchmark(b *testing.B) {
	type input struct {
		X float64
		Y float64
	}

	exprcls.RegisterSource[input]()

	// Expression to be evaluated using both approaches.
	testExpressionStr := "X+(6*Y)"

	// Compile the expression string into a callable Go function with expr-cls.
	exprClsProg, err := exprcls.Compile[input, float64](testExpressionStr)
	require.Nil(b, err)

	// Compile the same expression using the external library github.com/expr-lang/expr.
	exprProg, err := expr.Compile(testExpressionStr, expr.AsFloat64(), expr.Env(input{}))
	require.Nil(b, err)

	testExpressionParams := input{X: 3, Y: 5}
	// Ensure both methods yield the same result.
	val, _ := expr.Run(exprProg, testExpressionParams)
	require.Equal(b, exprClsProg(testExpressionParams), val)

	// Display details about the expression being tested.
	fmt.Printf("expression: %q\n", testExpressionStr)
	fmt.Printf("params:\n\tX=%v\n\tY=%v\n", testExpressionParams.X, testExpressionParams.Y)
	fmt.Printf("expr-cls result: %v\n", exprClsProg(testExpressionParams))
	fmt.Printf("expr result: %v\n\n", val)

	// Benchmark the expr-cls MVP build.
	b.Run("expr-cls", func(b *testing.B) {
		for b.Loop() {
			exprClsProg(testExpressionParams)
		}
	})

	// Benchmark the external library's compilation approach.
	b.Run("expr", func(b *testing.B) {
		for b.Loop() {
			expr.Run(exprProg, testExpressionParams)
		}
	})
}
