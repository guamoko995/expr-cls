package env

import (
	"fmt"
	"math"

	"github.com/guamoko995/expr-cls/builder/makebuilder"
	"github.com/guamoko995/expr-cls/internal/hash"
)

func (env *Env) SetDefaultConsts() {
	defConstViaBuilderNoErr("pi",
		makebuilder.Const(math.Pi),
	)
	defConstViaBuilderNoErr("phi",
		makebuilder.Const(math.Phi),
	)
}

func defConstViaBuilderNoErr(token string, builder ...any) {
	if err := DefConstViaBuilder(token, Global, builder...); err != nil {
		panic(err)
	}
}

func DefConstViaBuilder(token string, env *Env, builder ...any) error {
	if _, exist := env.Unares[token]; !exist {
		return fmt.Errorf("operator %q is not supported", token)
	}
	for i := range builder {

		env.Unares[token][hash.HashArgsByBuilder(builder[i])] = builder
	}

	return nil
}
