package env

import (
	"maps"

	"github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/env/registrators"
)

// RegisterVarType adds variable types to the environment registry.
// If the provided token matches any operator, an error is returned.
func (env *Enviroment) RegisterVarType(registrator ...registrators.VariableType) {
	for i := range registrator {
		env.VariableMakers[registrator[i].GetVarType()] = registrator[i]
		for srcTR := range env.Variables {
			maps.Copy(env.Variables[srcTR], registrator[i].MakeGetVarBuildersFromSrc(srcTR))
		}
	}
}

// RegisterVarType registers variable types globally using the default
// environment instance. This function delegates the actual work to the
// RegisterVarType method of the global DefaultEnv object.
func RegisterVarType(registrator ...registrators.VariableType) {
	DefaultEnv.RegisterVarType(registrator...)
}

// RegisterVarSource adds the variable source type to the environment registry.
// If the provided token matches any operator, an error is returned.
func (env *Enviroment) RegisterVarSource(registrator ...registrators.VarSourse) {
	for i := range registrator {
		env.Variables[registrator[i].GetSrcType()] = make(map[string]base.Builder)
		for _, varBuildersMaker := range env.VariableMakers {
			maps.Copy(env.Variables[registrator[i].GetSrcType()], varBuildersMaker.MakeGetVarBuildersFromSrc(registrator[i].GetSrcType()))
		}
	}
}

// RegisterVarTypes registers a variable source type globally using the default
// environment instance. This function delegates the actual work to the
// RegisterVarSourse method of the global DefaultEnv object.
func RegisterVarSource(registrator ...registrators.VarSourse) {
	DefaultEnv.RegisterVarSource(registrator...)
}
