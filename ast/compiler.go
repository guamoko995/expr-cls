package ast

import (
	"fmt"
	"reflect"

	"github.com/guamoko995/expr-cls/env"
	basepkg "github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/env/registrators"
)

func Compile[srcT, outT any, pSrcT *srcT](tree Node, e *env.Enviroment) (func(src srcT) outT, error) {
	var srcC srcT
	pSrc := &srcC

	if _, exist := e.Variables[reflect.TypeFor[srcT]()]; !exist {
		env.RegisterVarSources(registrators.NewVarSourse[srcT]())
	}

	fnAny, err := tree.Build(e, pSrc)
	if err != nil {
		return nil, err
	}

	fn, ok := fnAny.(basepkg.LazyFunc[outT])
	if !ok {
		return nil, fmt.Errorf("unexpectet type result: %T", fnAny)
	}

	return func(src srcT) outT {
		srcC = src
		return fn()
	}, nil
}
