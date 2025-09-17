package registrators

import (
	"reflect"

	base "github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/internal/hashsum"
)

type Const struct {
	builder base.Builder
}

func (cb Const) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return cb.builder.Build(args)
}

func (cb Const) GetOutType() reflect.Type {
	return cb.builder.GetOutType()
}
func (cb Const) GetInputTypesHashSum() hashsum.Inputs {
	return 0
}

// constant creates a lazy constructor based on a value.
type constant[T any] base.LazyFunc[T]

func (constant[T]) GetOutType() reflect.Type {
	return reflect.TypeFor[T]()
}
func (constant[T]) GetInputTypesHashSum() hashsum.Inputs {
	return 0
}

// Build forms a lazy constructor for a constant.
func (c constant[T]) Build([]base.GenericLazyFunc) base.GenericLazyFunc {
	return c
}

// NewConst turns a value into a lazy constructor.
func NewConst[T any](val T) Const {
	return Const{
		constant[T](
			base.LazyFunc[T](
				func() T {
					return val
				},
			),
		),
	}
}
