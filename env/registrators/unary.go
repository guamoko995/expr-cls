package registrators

import (
	"reflect"

	base "github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/internal/hashsum"
)

type Unary struct {
	builder base.Builder
}

func (op Unary) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return op.builder.Build(args)
}

func (op Unary) GetOutType() reflect.Type {
	return op.builder.GetOutType()
}
func (op Unary) GetInputTypesHashSum() hashsum.Inputs {
	return op.builder.GetInputTypesHashSum()
}

// NewUnary turns a function with one input parameter into a lazy
// constructor.
func NewUnary[in1T, outT any](
	src func(in1T) outT,
) Unary {
	return Unary{funcIn1[in1T, outT](func(
		lazyF base.LazyFunc[in1T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF())
		}
	})}
}
