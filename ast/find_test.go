package ast_test

import (
	"testing"

	"github.com/guamoko995/expr-cls/internal/testify/require"

	"github.com/guamoko995/expr-cls/ast"
)

func TestFind(t *testing.T) {
	left := &ast.IdentifierNode{
		Value: "a",
	}
	var root ast.Node = &ast.BinaryNode{
		Operator: "+",
		Left:     left,
		Right: &ast.IdentifierNode{
			Value: "b",
		},
	}

	x := ast.Find(root, func(node ast.Node) bool {
		if n, ok := node.(*ast.IdentifierNode); ok {
			return n.Value == "a"
		}
		return false
	})

	require.Equal(t, left, x)
}
