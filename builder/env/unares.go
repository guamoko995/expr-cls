package env

import (
	"fmt"

	"github.com/guamoko995/expr-cls/builder/base"
	"github.com/guamoko995/expr-cls/internal/hash"
)

func (env *Env) RegDefaultUnares() {
	defUnaryViaBuilderNoErr("not",
		base.MakeFuncIn1Builder(func(a bool) bool { return !a }),
	)

	defUnaryViaBuilderNoErr("!",
		base.MakeFuncIn1Builder(func(a bool) bool { return !a }),
	)

	defUnaryViaBuilderNoErr("-",
		base.MakeFuncIn1Builder(func(a int) int { return -a }),
		base.MakeFuncIn1Builder(func(a float64) float64 { return -a }),
	)
}

func defUnaryViaBuilderNoErr(token string, builder ...base.Builder) {
	if err := DefUnaryViaBuilder(token, Global, builder...); err != nil {
		panic(err)
	}
}

func DefUnaryViaBuilder(token string, env *Env, builder ...base.Builder) error {
	if _, exist := env.UnaryBuilders[token]; !exist {
		return fmt.Errorf("operator %q is not supported", token)
	}
	for i := range builder {
		env.UnaryBuilders[token][hash.HashArgsByBuilder(builder[i])] = builder[i]
	}

	return nil
}
