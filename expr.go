package exprcls

import (
	"github.com/guamoko995/expr-cls/ast"
	"github.com/guamoko995/expr-cls/conf"
	"github.com/guamoko995/expr-cls/parser"
)

// Option for configuring config.
type Option func(c *conf.Config)

// MaxNodes sets the maximum number of nodes allowed in the expression.
// By default, the maximum number of nodes is conf.DefaultMaxNodes.
// If MaxNodes is set to 0, the node budget check is disabled.
func MaxNodes(n uint) Option {
	return func(c *conf.Config) {
		c.MaxNodes = n
	}
}

func WithEnv(Env Env) Option {
	return func(c *conf.Config) {
		c.WithEnv(Env.internalEnv.Enviroment)
	}
}

// Compile parses and compiles given input expression to bytecode program.
func Compile[srcT, outT any](input string, ops ...Option) (func(srcT) outT, error) {
	config := conf.New()
	for _, op := range ops {
		op(config)
	}

	tree, err := parser.ParseWithConfig(input, config)
	if err != nil {
		return nil, err
	}

	res, err := ast.Compile[srcT, outT](tree.Node, config.BuildStageEnvironment)
	if err != nil {
		return nil, err
	}

	return res, nil
}
