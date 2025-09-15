package env

import (
	"reflect"

	"github.com/guamoko995/expr-cls/builder/base"
	"github.com/guamoko995/expr-cls/parser/operator"

	"github.com/guamoko995/expr-cls/internal/hash"
)

func init() {
	Global.RegDefaultEnv()
}

var Global = New()

type Env struct {
	UnaryBuilders    map[string]map[hash.Args]base.Builder
	BinaryBuilders   map[string]map[hash.Args]base.Builder
	FunctionBuilders map[string]map[hash.Args]base.Builder
	Consts           map[string]base.GenericLazyFunc
	BuilderMakers    map[reflect.Type]BuildersMaker
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

	env.BuilderMakers = make(map[reflect.Type]BuildersMaker)
	env.VarBuilders = make(map[reflect.Type]map[string]base.Builder)

	return &env
}

func (e *Env) RegDefaultEnv() {
	e.RegDefaultUnares()
	e.RegDefaultBinares()
	e.RegDefaultFuncs()
	e.RegDefaultConsts()
	e.RegDefaultVarTypes()
}
