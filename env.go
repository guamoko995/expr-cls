package exprcls

import (
	"github.com/guamoko995/expr-cls/env"
	"github.com/guamoko995/expr-cls/env/registrators"
)

type internalEnv struct {
	*env.Enviroment
}

type Env struct {
	internalEnv
}

func NewEnv() Env {
	return Env{internalEnv{env.New()}}
}

func RegisterUnary(token string, registrator ...registrators.Unary) error {
	return env.RegisterUnary(token, registrator...)
}

func RegisterBinary(token string, registrator ...registrators.Binary) error {
	return env.RegisterBinary(token, registrator...)
}

func RegisterFunc(token string, registrator ...registrators.Func) error {
	return env.RegisterFunc(token, registrator...)
}

func RegisterConst(token string, registrator ...registrators.Const) error {
	return env.RegisterConst(token, registrator...)
}

func RegisterVarType(registrator ...registrators.VariableType) {
	env.RegisterVarType(registrator...)
}

func RegisterVarSource(registrator ...registrators.VarSourse) {
	env.RegisterVarSource(registrator...)
}
