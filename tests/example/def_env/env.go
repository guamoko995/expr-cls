package defenv

import "github.com/guamoko995/expr-cls/env"

func init() {
	DefEnv(env.DefaultEnv)
}

// DefEnv defines the environment.
func DefEnv(e *env.Enviroment) {
	RegisterUnares(e)
	RegisterBinares(e)
	RegisterFuncs(e)
	RegisterConsts(e)
	RegisterVarTypes(e)
}
