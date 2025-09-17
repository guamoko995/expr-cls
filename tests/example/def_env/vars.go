package defenv

import (
	"github.com/guamoko995/expr-cls/env"
	"github.com/guamoko995/expr-cls/env/registrators"
)

// RegisterVarTypes registers var types in the environment.
func RegisterVarTypes(env *env.Enviroment) {
	env.RegisterVarType(
		registrators.NewVar[int](),
		registrators.NewVar[float64](),
		registrators.NewVar[string](),
		registrators.NewVar[bool](),
	)
}
