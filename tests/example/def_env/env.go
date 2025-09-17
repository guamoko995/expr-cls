package defenv

import "github.com/guamoko995/expr-cls/env"

func init() {
	Regster(env.DefaultEnv)
}

func Regster(e *env.Enviroment) {
	RegisterUnares(e)
	RegisterBinares(e)
	RegisterFuncs(e)
	RegisterConsts(e)
	RegisterVarTypes(e)
}
