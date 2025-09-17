package env

import (
	"reflect"

	"github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/env/registrators"
	"github.com/guamoko995/expr-cls/parser/operator"

	"github.com/guamoko995/expr-cls/internal/hashsum"
)

var DefaultEnv = New()

type Enviroment struct {
	Unary          map[string]map[hashsum.Inputs]base.Builder
	Binary         map[string]map[hashsum.Inputs]base.Builder
	Func           map[string]map[hashsum.Inputs]base.Builder
	Const          map[string]registrators.Const
	VariableMakers map[reflect.Type]registrators.VariableType
	Variables      map[reflect.Type]map[string]base.Builder
}

func New() *Enviroment {
	var env Enviroment

	env.Unary = make(map[string]map[hashsum.Inputs]base.Builder, len(operator.Unary))
	for key := range operator.Unary {
		env.Unary[key] = make(map[hashsum.Inputs]base.Builder)
	}

	env.Binary = make(map[string]map[hashsum.Inputs]base.Builder, len(operator.Binary))
	for key := range operator.Binary {
		env.Binary[key] = make(map[hashsum.Inputs]base.Builder)
	}

	env.Func = make(map[string]map[hashsum.Inputs]base.Builder)

	env.Const = make(map[string]registrators.Const)

	env.VariableMakers = make(map[reflect.Type]registrators.VariableType)
	env.Variables = make(map[reflect.Type]map[string]base.Builder)

	return &env
}
