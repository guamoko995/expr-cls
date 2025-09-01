package main

import (
	"os"
	"testing"

	"github.com/guamoko995/expr-cls"
	"github.com/guamoko995/expr-cls/internal/testify/require"
)

var examples []CodeBlock

func init() {
	b, err := os.ReadFile("../../testdata/examples.md")
	if err != nil {
		panic(err)
	}
	examples = extractCodeBlocksWithTitle(string(b))
}

func TestExamples(t *testing.T) {
	for _, code := range examples {
		code := code
		t.Run(code.Title, func(t *testing.T) {
			program, err := expr.Compile(code.Content, expr.Env(nil))
			require.NoError(t, err)

			_, err = expr.Run(program, nil)
			require.NoError(t, err)
		})
	}
}
