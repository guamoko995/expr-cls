package env

import (
	"fmt"

	"github.com/guamoko995/expr-cls/builder/makebuilder"
	"github.com/guamoko995/expr-cls/internal/hash"
)

func (env *Env) SetDefaultUnares() {
	defUnaryViaBuilderNoErr("not",
		makebuilder.FuncIn1(func(a bool) bool { return !a }),
	)

	defUnaryViaBuilderNoErr("!",
		makebuilder.FuncIn1(func(a bool) bool { return !a }),
	)

	defUnaryViaBuilderNoErr("-",
		makebuilder.FuncIn1(func(a int64) int64 { return -a }),
		makebuilder.FuncIn1(func(a float64) float64 { return -a }),
	)
}

func defUnaryViaBuilderNoErr(token string, builder ...any) {
	if err := DefConstViaBuilder(token, Global, builder...); err != nil {
		panic(err)
	}
}

func DefUnaryViaBuilder(token string, env *Env, builder ...any) error {
	if _, exist := env.Unares[token]; !exist {
		return fmt.Errorf("operator %q is not supported", token)
	}
	for i := range builder {

		env.Unares[token][hash.HashArgsByBuilder(builder[i])] = builder
	}

	return nil
}
