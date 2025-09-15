package builder

import (
	"fmt"

	"github.com/guamoko995/expr-cls/ast"
	"github.com/guamoko995/expr-cls/builder/base"
	"github.com/guamoko995/expr-cls/builder/env"
)

func Build[srcT, outT any, pSrcT *srcT](tree ast.Node, env *env.Env) (func(src srcT) outT, error) {
	var srcC srcT
	var pSrc pSrcT
	pSrc = &srcC
	fnAny, err := tree.Build(env, pSrc)
	if err != nil {
		return nil, err
	}

	fn, ok := fnAny.(base.LazyFunc[outT])
	if !ok {
		return nil, fmt.Errorf("unexpectet type result: %T", fnAny)
	}

	return func(src srcT) outT {
		srcC = src
		return fn()
	}, nil
}
