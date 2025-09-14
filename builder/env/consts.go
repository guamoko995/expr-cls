package env

import (
	"fmt"
	"math"
)

func (env *Env) SetDefaultConsts() {
	defConst("pi", math.Pi)
	defConst("phi", math.Phi)
}

func defConst[T any](token string, val T) {
	if err := DefConst(token, Global, val); err != nil {
		panic(err)
	}
}

func DefConst[T any](token string, env *Env, val T) error {
	if _, exist := env.UnaryBuilders[token]; exist {
		return fmt.Errorf("token %q is reserved for operators", token)
	}

	if _, exist := env.BinaryBuilders[token]; exist {
		return fmt.Errorf("token %q is reserved for operators", token)
	}

	if _, exist := env.FunctionBuilders[token]; exist {
		return fmt.Errorf("token %q is reserved by function", token)
	}

	env.Consts[token] = func() T {
		return val
	}

	return nil
}
