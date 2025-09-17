package defenv

import (
	"math"

	"github.com/guamoko995/expr-cls/env"
	"github.com/guamoko995/expr-cls/env/registrators"
)

// RegisterConsts registers constant in the environment.
func RegisterConsts(env *env.Enviroment) {
	env.RegisterConstNoErr("pi", registrators.NewConst(math.Pi))
	env.RegisterConstNoErr("phi", registrators.NewConst(math.Phi))
}
