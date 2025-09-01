//go:build expr_debug
// +build expr_debug

package vm_test

import (
	"testing"

	vm "github.com/guamoko995/expr-cls/TO_REMOVE_vm"
	"github.com/guamoko995/expr-cls/internal/testify/require"

	"github.com/guamoko995/expr-cls/compiler"
	"github.com/guamoko995/expr-cls/parser"
)

func TestDebugger(t *testing.T) {
	input := `[1, 2]`

	node, err := parser.Parse(input)
	require.NoError(t, err)

	program, err := compiler.Compile(node, nil)
	require.NoError(t, err)

	debug := vm.Debug()
	go func() {
		debug.Step()
		debug.Step()
		debug.Step()
		debug.Step()
	}()
	go func() {
		for range debug.Position() {
		}
	}()

	_, err = debug.Run(program, nil)
	require.NoError(t, err)
	require.Len(t, debug.Stack, 0)
	require.Nil(t, debug.Scopes)
}
