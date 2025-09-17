package defenv

import (
	"github.com/guamoko995/expr-cls/env"
	"github.com/guamoko995/expr-cls/env/registrators"
)

// RegisterUnares registers unary operators in the environment.
func RegisterUnares(env *env.Enviroment) {
	env.RegisterUnaryNoErr("not",
		registrators.NewUnary(func(a bool) bool { return !a }),
	)

	env.RegisterUnaryNoErr("!",
		registrators.NewUnary(func(a bool) bool { return !a }),
	)

	env.RegisterUnaryNoErr("-",
		registrators.NewUnary(func(a int) int { return -a }),
		registrators.NewUnary(func(a float64) float64 { return -a }),
	)
}
