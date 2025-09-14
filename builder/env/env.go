package env

import (
	"reflect"

	"github.com/guamoko995/expr-cls/builder/base"
	"github.com/guamoko995/expr-cls/parser/operator"

	"github.com/guamoko995/expr-cls/internal/hash"
)

var Global = New()

type Env struct {
	UnaryBuilders    map[string]map[hash.Args]base.Builder
	BinaryBuilders   map[string]map[hash.Args]base.Builder
	FunctionBuilders map[string]map[hash.Args]base.Builder
	Consts           map[string]base.GenericLazyFunc
	VarBuilderMakers map[reflect.Type]anyGetVarBuildersMaker
	VarBuilders      map[reflect.Type]map[string]base.Builder
}

func New() *Env {
	var env Env

	env.UnaryBuilders = make(map[string]map[hash.Args]base.Builder, len(operator.Unary))
	for key := range operator.Unary {
		env.UnaryBuilders[key] = make(map[hash.Args]base.Builder)
	}

	env.BinaryBuilders = make(map[string]map[hash.Args]base.Builder, len(operator.Binary))
	for key := range operator.Binary {
		env.BinaryBuilders[key] = make(map[hash.Args]base.Builder)
	}

	env.FunctionBuilders = make(map[string]map[hash.Args]base.Builder)

	env.Consts = make(map[string]base.GenericLazyFunc)

	env.VarBuilderMakers = make(map[reflect.Type]anyGetVarBuildersMaker)
	env.VarBuilders = make(map[reflect.Type]map[string]base.Builder)

	return &env
}

func (e *Env) SetDefaultEnv() {
	e.SetDefaultUnares()
	e.SetDefaultBinares()
	e.SetDefaultFuncs()
	e.SetDefaultConsts()
	e.SetDefaultVarTypes()
}
