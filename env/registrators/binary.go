package registrators

import (
	"reflect"

	base "github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/internal/hashsum"
)

type Binary struct {
	builder base.Builder
}

func (op Binary) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return op.builder.Build(args)
}

func (op Binary) GetOutType() reflect.Type {
	return op.builder.GetOutType()
}
func (op Binary) GetInputTypesHashSum() hashsum.Inputs {
	return op.builder.GetInputTypesHashSum()
}

// NewBinary turns a function with two input parameter into a lazy constructor.
func NewBinary[in1T, in2T, outT any](
	src func(in1T, in2T) outT,
) Binary {
	return Binary{funcIn2[in1T, in2T, outT](func(
		lazyF1 base.LazyFunc[in1T],
		lazyF2 base.LazyFunc[in2T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2())
		}
	})}
}
