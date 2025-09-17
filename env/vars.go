package env

import (
	"maps"

	"github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/env/registrators"
)

func (env *Enviroment) RegisterVarTypes(registrator ...registrators.VariableType) {
	for i := range registrator {
		env.VariableMakers[registrator[i].GetVarType()] = registrator[i]
		for srcTR := range env.Variables {
			maps.Copy(env.Variables[srcTR], registrator[i].MakeGetVarBuildersFromSrc(srcTR))
		}
	}
}

func RegisterVarTypes(registrator ...registrators.VariableType) {
	DefaultEnv.RegisterVarTypes(registrator...)
}

func (env *Enviroment) RegisterVarSources(registrator ...registrators.VarSourse) {
	for i := range registrator {
		env.Variables[registrator[i].GetSrcType()] = make(map[string]base.Builder)
		for _, varBuildersMaker := range env.VariableMakers {
			maps.Copy(env.Variables[registrator[i].GetSrcType()], varBuildersMaker.MakeGetVarBuildersFromSrc(registrator[i].GetSrcType()))
		}
	}
}

func RegisterVarSources(registrator ...registrators.VarSourse) {
	DefaultEnv.RegisterVarSources(registrator...)
}
