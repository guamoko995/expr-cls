package defenv

import (
	"math"

	"github.com/guamoko995/expr-cls/env"
	"github.com/guamoko995/expr-cls/env/registrators"
)

// RegisterConsts registers default constant in the current environment.
func RegisterConsts(env *env.Enviroment) {
	env.RegisterConstNoErr("pi", registrators.NewConst(math.Pi))
	env.RegisterConstNoErr("phi", registrators.NewConst(math.Phi))
}
