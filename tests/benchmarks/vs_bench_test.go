package poc

import (
	"fmt"
	"testing"

	"github.com/expr-lang/expr"
	"github.com/google/cel-go/cel"
	exprcls "github.com/guamoko995/expr-cls"

	// Using the example environment
	_ "github.com/guamoko995/expr-cls/tests/example/def_env"
	"github.com/stretchr/testify/require"
)

// The test compares the performance of two approaches to expression compilation:
// 1. Using github.com/guamoko995/expr-cls.
// 2. Using github.com/expr-lang/expr.
// 3. Using github.com/google/cel-go.
func BenchmarkCompile(b *testing.B) {
	type input struct {
		X float64
		Y float64
	}

	// Expression to be evaluated using both approaches.
	testExpressionStr := "X+(6.0*Y)"

	_, err := exprcls.Compile[input, float64](testExpressionStr)
	require.Nil(b, err)

	// Benchmark the expr-cls.
	b.Run("expr-cls", func(b *testing.B) {
		for b.Loop() {
			// Compile the expression string into a callable Go function with expr-cls.
			exprcls.Compile[input, float64](testExpressionStr)
		}
	})

	_, err = expr.Compile(testExpressionStr, expr.AsFloat64(), expr.Env(input{}))
	require.Nil(b, err)

	// Benchmark the expr.
	b.Run("expr", func(b *testing.B) {
		for b.Loop() {
			// Compile the same expression using the external library github.com/expr-lang/expr.
			expr.Compile(testExpressionStr, expr.AsFloat64(), expr.Env(input{}))
		}
	})

	env, err := cel.NewEnv(
		cel.Variable("X", cel.DoubleType),
		cel.Variable("Y", cel.DoubleType),
	)
	require.Nil(b, err)

	ast, issues := env.Compile(testExpressionStr)
	if issues != nil {
		require.Nil(b, err)
	}
	_, err = env.Program(ast)
	require.Nil(b, err)

	// Benchmark the cel-go.
	b.Run("cel-go", func(b *testing.B) {
		for b.Loop() {
			ast, _ := env.Compile(testExpressionStr)
			env.Program(ast)
		}
	})

}

// The test compares the performance of two approaches to expression calculation:
// 1. Using github.com/guamoko995/expr-cls.
// 2. Using github.com/expr-lang/expr.
func BenchmarkСalc(b *testing.B) {
	type input struct {
		X float64
		Y float64
	}

	// Expression to be evaluated using both approaches.
	testExpressionStr := "X+(6.0*Y)"

	// Compile the expression string into a callable Go function with expr-cls.
	exprClsProg, err := exprcls.Compile[input, float64](testExpressionStr)
	require.Nil(b, err)

	// Compile the same expression using the external library github.com/expr-lang/expr.
	exprProg, err := expr.Compile(testExpressionStr, expr.AsFloat64(), expr.Env(input{}))
	require.Nil(b, err)

	// Compile the same expression using the external library "github.com/google/cel-go/cel".
	env, err := cel.NewEnv(
		cel.Variable("X", cel.DoubleType),
		cel.Variable("Y", cel.DoubleType),
	)
	require.Nil(b, err)
	ast, issues := env.Compile(testExpressionStr)
	if issues != nil {
		require.Nil(b, err)
	}
	celProg, err := env.Program(ast)
	require.Nil(b, err)
	testExpressionParams := input{X: 3, Y: 5}

	// Ensure both methods yield the same result.
	valExpr, err := expr.Run(exprProg, testExpressionParams)
	require.Nil(b, err)
	require.Equal(b, exprClsProg(testExpressionParams), valExpr)

	valCel, _, err := celProg.Eval(map[string]interface{}{"X": 3.0, "Y": 5.0})
	require.Nil(b, err)
	require.Equal(b, exprClsProg(testExpressionParams), valCel.Value().(float64))

	// Display details about the expression being tested.
	fmt.Printf("expression: %q\n", testExpressionStr)
	fmt.Printf("params:\n\tX=%v\n\tY=%v\n", testExpressionParams.X, testExpressionParams.Y)
	fmt.Printf("expr-cls result: %v\n", exprClsProg(testExpressionParams))
	fmt.Printf("expr result: %v\n", valExpr)
	fmt.Printf("cel-go result: %v\n\n", valCel)

	// Benchmark the expr-cls.
	b.Run("expr-cls", func(b *testing.B) {
		for b.Loop() {
			exprClsProg(testExpressionParams)
		}
	})

	// Benchmark the expr.
	b.Run("expr", func(b *testing.B) {
		for b.Loop() {
			expr.Run(exprProg, testExpressionParams)
		}
	})

	// Benchmark the cel-go.
	b.Run("cel-go", func(b *testing.B) {
		for b.Loop() {
			celProg.Eval(map[string]interface{}{"X": 3.0, "Y": 5.0})
		}
	})
}
