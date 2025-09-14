package builder

import (
	"github.com/guamoko995/expr-cls/ast"
	"github.com/guamoko995/expr-cls/parser"
)

type fabric struct {
	builders map[*ast.Node]any
}

type buildNode any

func (fab *fabric) Visit(node *ast.Node) {

}

func Fabr[resultT, envT any](tree *parser.Tree) (func(envT) resultT, error) {

	return nil, nil
}
