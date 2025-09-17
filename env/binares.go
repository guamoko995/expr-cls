package env

import (
	"fmt"

	"github.com/guamoko995/expr-cls/env/registrators"
)

func (env *Enviroment) RegisterBinaryNoErr(token string, registrator ...registrators.Binary) {
	if err := env.RegisterBinary(token, registrator...); err != nil {
		panic(err)
	}
}

// RegisterBinary adds a set of builders for a given binary operator into the
// environment's registry. This method allows associating different
// implementations of a binary operation based on input types. If the provided
// token does not correspond to any known binary operator, returns an error.
func (env *Enviroment) RegisterBinary(token string, registrator ...registrators.Binary) error {
	if _, exist := env.Binary[token]; !exist {
		return fmt.Errorf("operator %q is not supported", token)
	}

	for i := range registrator {
		env.Binary[token][registrator[i].GetInputTypesHashSum()] = registrator[i]
	}
	return nil
}

// RegisterBinary registers a set of builders for a given binary operator
// globally using the default environment instance. This function delegates
// the actual work to the RegisterBinary method of the global DefaultEnv object.
func RegisterBinary(token string, registrator ...registrators.Binary) error {
	return DefaultEnv.RegisterBinary(token, registrator...)
}
