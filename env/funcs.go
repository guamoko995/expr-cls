package env

import (
	"errors"
	"fmt"

	"github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/internal/hashsum"

	"github.com/guamoko995/expr-cls/env/registrators"
)

// RegisterFuncNoErr calls RegisterFunc and panics on error.
func (env *Enviroment) RegisterFuncNoErr(token string, registrator ...registrators.Func) {
	for i := range registrator {
		if err := env.RegisterFunc(token, registrator...); err != nil {
			panic(fmt.Errorf("failed def %d (starting from 0) builder: %w", i, err))
		}
	}
}

// RegisterFunc adds a set of collectors for the given function to the
// environment registry. This method allows you to chain together different
// implementations of a function based on input types. If the provided token
// matches any operator, an error is returned.
func (env *Enviroment) RegisterFunc(token string, registrator ...registrators.Func) error {
	if _, exist := env.Unary[token]; exist {
		return errors.New("token is reserved for operators")
	}

	if _, exist := env.Binary[token]; exist {
		return errors.New("token is reserved for operators")
	}

	delete(env.Const, token)

	if _, exist := env.Func[token]; !exist {
		env.Func[token] = make(map[hashsum.Inputs]base.Builder)
	}

	for i := range registrator {
		env.Func[token][registrator[i].GetInputTypesHashSum()] = registrator[i]
	}

	return nil
}

// RegisterFunc registers a set of builders for a given function
// globally using the default environment instance. This function delegates
// the actual work to the RegisterFunc method of the global DefaultEnv object.
func RegisterFunc(token string, registrator ...registrators.Func) error {
	return DefaultEnv.RegisterFunc(token, registrator...)
}
