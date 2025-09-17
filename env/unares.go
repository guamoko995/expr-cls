package env

import (
	"fmt"

	"github.com/guamoko995/expr-cls/env/registrators"
)

func (env *Enviroment) RegisterUnaryNoErr(token string, registrator ...registrators.Unary) {
	if err := RegisterUnary(token, registrator...); err != nil {
		panic(err)
	}
}

// RegisterUnary adds a set of builders for a given unary operator into the
// environment's registry. This method allows associating different
// implementations of a unary operation based on input types. If the provided
// token does not correspond to any known unary operator, returns an error.
func (env *Enviroment) RegisterUnary(token string, registrator ...registrators.Unary) error {
	if _, exist := env.Unary[token]; !exist {
		return fmt.Errorf("operator %q is not supported", token)
	}
	for i := range registrator {
		env.Unary[token][registrator[i].GetInputTypesHashSum()] = registrator[i]
	}

	return nil
}

// RegisterUnary registers a set of builders for a given unary operator
// globally using the default environment instance. This function delegates
// the actual work to the RegisterUnary method of the global DefaultEnv object.
func RegisterUnary(token string, registrator ...registrators.Unary) error {
	return DefaultEnv.RegisterUnary(token, registrator...)
}
