package defenv

import (
	"github.com/guamoko995/expr-cls/env"
	"github.com/guamoko995/expr-cls/env/registrators"
)

func RegisterVarTypes(env *env.Enviroment) {
	env.RegisterVarTypes(
		registrators.NewVar[int](),
		registrators.NewVar[float64](),
		registrators.NewVar[string](),
		registrators.NewVar[bool](),
	)
}
