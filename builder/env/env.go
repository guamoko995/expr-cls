package env

import (
	"github.com/guamoko995/expr-cls/parser/operator"

	"github.com/guamoko995/expr-cls/internal/hash"
)

var Global = New()

type Env struct {
	Unares    map[string]map[hash.Args]any
	Binares   map[string]map[hash.Args]any
	Functions map[string]map[hash.Args]any
	Consts    map[string]any
}

func New() *Env {
	var env Env

	env.Unares = make(map[string]map[hash.Args]any, len(operator.Unary))
	for key := range operator.Unary {
		env.Unares[key] = make(map[hash.Args]any)
	}

	env.Binares = make(map[string]map[hash.Args]any, len(operator.Binary))
	for key := range operator.Binary {
		env.Binares[key] = make(map[hash.Args]any)
	}

	env.Functions = make(map[string]map[hash.Args]any)
	env.Consts = make(map[string]any)

	return &env
}
