package env

import (
	"fmt"

	"github.com/guamoko995/expr-cls/env/registrators"
)

// RegisterConstNoErr calls RegisterConst and panics on error.
func (env *Enviroment) RegisterConstNoErr(token string, registrator ...registrators.Const) {
	if err := env.RegisterConst(token, registrator...); err != nil {
		panic(err)
	}
}

// RegisterConst adds a set of collectors for the given constant
// to the environment registry. If the provided token matches any operator,
// an error is returned.
func (env *Enviroment) RegisterConst(token string, registrator ...registrators.Const) error {
	if _, exist := env.Unary[token]; exist {
		return fmt.Errorf("token %q is reserved for operators", token)
	}

	if _, exist := env.Binary[token]; exist {
		return fmt.Errorf("token %q is reserved for operators", token)
	}

	delete(env.Func, token)

	for i := range registrator {
		env.Const[token] = registrator[i]
	}

	return nil
}

// RegisterConst registers a set of builders for a given constant globally
// using the default environment instance. This function delegates the actual
// work to the RegisterConst method of the global DefaultEnv object.
func RegisterConst(token string, registrator ...registrators.Const) error {
	return DefaultEnv.RegisterConst(token, registrator...)
}
