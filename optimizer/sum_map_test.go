package optimizer_test

import (
	"testing"

	"github.com/guamoko995/expr-cls/internal/testify/assert"
	"github.com/guamoko995/expr-cls/internal/testify/require"

	"github.com/guamoko995/expr-cls/ast"
	"github.com/guamoko995/expr-cls/optimizer"
	"github.com/guamoko995/expr-cls/parser"
)

func TestOptimize_sum_map(t *testing.T) {
	tree, err := parser.Parse(`sum(map(users, {.Age}))`)
	require.NoError(t, err)

	err = optimizer.Optimize(&tree.Node, nil)
	require.NoError(t, err)

	expected := &ast.BuiltinNode{
		Name: "sum",
		Arguments: []ast.Node{
			&ast.IdentifierNode{Value: "users"},
			&ast.PredicateNode{
				Node: &ast.MemberNode{
					Node:     &ast.PointerNode{},
					Property: &ast.StringNode{Value: "Age"},
				},
			},
		},
	}

	assert.Equal(t, ast.Dump(expected), ast.Dump(tree.Node))
}
