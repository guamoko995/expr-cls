package registrators

import (
	"reflect"

	"github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/internal/hashsum"
)

type Func struct {
	builder base.Builder
}

func (fb Func) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return fb.builder.Build(args)
}

func (fb Func) GetOutType() reflect.Type {
	return fb.builder.GetOutType()
}
func (fb Func) GetInputTypesHashSum() hashsum.Inputs {
	return fb.builder.GetInputTypesHashSum()
}

// funcIn1 creates a lazy constructor based on a function with one input
// parameter.
type funcIn1[in1T, outT any] func(base.LazyFunc[in1T]) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with one input parameter.
func (builder funcIn1[in1T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(args[0].(base.LazyFunc[in1T]))
}
func (builder funcIn1[in1T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn1[in1T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn1 turns a function with one input parameter into a lazy
// constructor.
func NewFuncIn1[in1T, outT any](
	src func(in1T) outT,
) Func {
	return Func{funcIn1[in1T, outT](func(
		lazyF base.LazyFunc[in1T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF())
		}
	})}
}

// funcIn2 creates a lazy constructor based on a function with two input
// parameters.
type funcIn2[in1T, in2T, outT any] func(base.LazyFunc[in1T], base.LazyFunc[in2T]) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with two input parameters.
func (builder funcIn2[in1T, in2T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(
		args[0].(base.LazyFunc[in1T]),
		args[1].(base.LazyFunc[in2T]),
	)
}
func (builder funcIn2[in1T, in2T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn2[in1T, in2T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn2 turns a function with two input parameters into a lazy
// constructor.
func NewFuncIn2[in1T, in2T, outT any](
	src func(in1T, in2T) outT,
) Func {
	return Func{funcIn2[in1T, in2T, outT](func(
		lazyF1 base.LazyFunc[in1T],
		lazyF2 base.LazyFunc[in2T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2())
		}
	})}
}

// funcIn3 creates a lazy constructor based on a function with three input
// parameters.
type funcIn3[in1T, in2T, in3T, outT any] func(base.LazyFunc[in1T], base.LazyFunc[in2T], base.LazyFunc[in3T]) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with three input parameters.
func (builder funcIn3[in1T, in2T, in3T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(
		args[0].(base.LazyFunc[in1T]),
		args[1].(base.LazyFunc[in2T]),
		args[2].(base.LazyFunc[in3T]),
	)
}
func (builder funcIn3[in1T, in2T, in3T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn3[in1T, in2T, in3T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn3 turns a function with three input parameters into a lazy
// constructor.
func NewFuncIn3[in1T, in2T, in3T, outT any](
	src func(in1T, in2T, in3T) outT,
) Func {
	return Func{funcIn3[in1T, in2T, in3T, outT](func(
		lazyF1 base.LazyFunc[in1T],
		lazyF2 base.LazyFunc[in2T],
		lazyF3 base.LazyFunc[in3T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3())
		}
	})}
}

// funcIn4 creates a lazy constructor based on a function with four input
// parameters.
type funcIn4[in1T, in2T, in3T, in4T, outT any] func(
	base.LazyFunc[in1T],
	base.LazyFunc[in2T],
	base.LazyFunc[in3T],
	base.LazyFunc[in4T],
) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with four input parameters.
func (builder funcIn4[in1T, in2T, in3T, in4T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(
		args[0].(base.LazyFunc[in1T]),
		args[1].(base.LazyFunc[in2T]),
		args[2].(base.LazyFunc[in3T]),
		args[3].(base.LazyFunc[in4T]),
	)
}
func (builder funcIn4[in1T, in2T, in3T, in4T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn4[in1T, in2T, in3T, in4T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn4 turns a function with four input parameters into a lazy
// constructor.
func NewFuncIn4[in1T, in2T, in3T, in4T, outT any](
	src func(in1T, in2T, in3T, in4T) outT,
) Func {
	return Func{funcIn4[in1T, in2T, in3T, in4T, outT](func(
		lazyF1 base.LazyFunc[in1T],
		lazyF2 base.LazyFunc[in2T],
		lazyF3 base.LazyFunc[in3T],
		lazyF4 base.LazyFunc[in4T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4())
		}
	})}
}

// funcIn5 creates a lazy constructor based on a function with five input
// parameters.
type funcIn5[in1T, in2T, in3T, in4T, in5T, outT any] func(
	base.LazyFunc[in1T],
	base.LazyFunc[in2T],
	base.LazyFunc[in3T],
	base.LazyFunc[in4T],
	base.LazyFunc[in5T],
) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with five input parameters.
func (builder funcIn5[in1T, in2T, in3T, in4T, in5T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(
		args[0].(base.LazyFunc[in1T]),
		args[1].(base.LazyFunc[in2T]),
		args[2].(base.LazyFunc[in3T]),
		args[3].(base.LazyFunc[in4T]),
		args[4].(base.LazyFunc[in5T]),
	)
}
func (builder funcIn5[in1T, in2T, in3T, in4T, in5T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn5[in1T, in2T, in3T, in4T, in5T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn5 turns a function with five input parameters into a lazy
// constructor.
func NewFuncIn5[in1T, in2T, in3T, in4T, in5T, outT any](
	src func(in1T, in2T, in3T, in4T, in5T) outT,
) Func {
	return Func{funcIn5[in1T, in2T, in3T, in4T, in5T, outT](func(
		lazyF1 base.LazyFunc[in1T],
		lazyF2 base.LazyFunc[in2T],
		lazyF3 base.LazyFunc[in3T],
		lazyF4 base.LazyFunc[in4T],
		lazyF5 base.LazyFunc[in5T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5())
		}
	})}
}

// funcIn6 creates a lazy constructor based on a function with six input
// parameters.
type funcIn6[in1T, in2T, in3T, in4T, in5T, in6T, outT any] func(
	base.LazyFunc[in1T],
	base.LazyFunc[in2T],
	base.LazyFunc[in3T],
	base.LazyFunc[in4T],
	base.LazyFunc[in5T],
	base.LazyFunc[in6T],
) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with six input parameters.
func (builder funcIn6[in1T, in2T, in3T, in4T, in5T, in6T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(
		args[0].(base.LazyFunc[in1T]),
		args[1].(base.LazyFunc[in2T]),
		args[2].(base.LazyFunc[in3T]),
		args[3].(base.LazyFunc[in4T]),
		args[4].(base.LazyFunc[in5T]),
		args[5].(base.LazyFunc[in6T]),
	)
}
func (builder funcIn6[in1T, in2T, in3T, in4T, in5T, in6T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn6[in1T, in2T, in3T, in4T, in5T, in6T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn6 turns a function with six input parameters into a lazy
// constructor.
func NewFuncIn6[in1T, in2T, in3T, in4T, in5T, in6T, outT any](
	src func(in1T, in2T, in3T, in4T, in5T, in6T) outT,
) Func {
	return Func{funcIn6[in1T, in2T, in3T, in4T, in5T, in6T, outT](func(
		lazyF1 base.LazyFunc[in1T],
		lazyF2 base.LazyFunc[in2T],
		lazyF3 base.LazyFunc[in3T],
		lazyF4 base.LazyFunc[in4T],
		lazyF5 base.LazyFunc[in5T],
		lazyF6 base.LazyFunc[in6T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6())
		}
	})}
}

// funcIn7 creates a lazy constructor based on a function with seven input
// parameters.
type funcIn7[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT any] func(
	base.LazyFunc[in1T],
	base.LazyFunc[in2T],
	base.LazyFunc[in3T],
	base.LazyFunc[in4T],
	base.LazyFunc[in5T],
	base.LazyFunc[in6T],
	base.LazyFunc[in7T],
) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with seven input parameters.
func (builder funcIn7[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(
		args[0].(base.LazyFunc[in1T]),
		args[1].(base.LazyFunc[in2T]),
		args[2].(base.LazyFunc[in3T]),
		args[3].(base.LazyFunc[in4T]),
		args[4].(base.LazyFunc[in5T]),
		args[5].(base.LazyFunc[in6T]),
		args[6].(base.LazyFunc[in7T]),
	)
}
func (builder funcIn7[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn7[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn7 turns a function with seven input parameters into a lazy
// constructor.
func NewFuncIn7[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT any](
	src func(in1T, in2T, in3T, in4T, in5T, in6T, in7T) outT,
) Func {
	return Func{funcIn7[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT](func(
		lazyF1 base.LazyFunc[in1T],
		lazyF2 base.LazyFunc[in2T],
		lazyF3 base.LazyFunc[in3T],
		lazyF4 base.LazyFunc[in4T],
		lazyF5 base.LazyFunc[in5T],
		lazyF6 base.LazyFunc[in6T],
		lazyF7 base.LazyFunc[in7T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6(), lazyF7())
		}
	})}
}

// funcIn8 creates a lazy constructor based on a function with eight input
// parameters.
type funcIn8[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT any] func(
	base.LazyFunc[in1T],
	base.LazyFunc[in2T],
	base.LazyFunc[in3T],
	base.LazyFunc[in4T],
	base.LazyFunc[in5T],
	base.LazyFunc[in6T],
	base.LazyFunc[in7T],
	base.LazyFunc[in8T],
) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with eight input parameters.
func (builder funcIn8[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(
		args[0].(base.LazyFunc[in1T]),
		args[1].(base.LazyFunc[in2T]),
		args[2].(base.LazyFunc[in3T]),
		args[3].(base.LazyFunc[in4T]),
		args[4].(base.LazyFunc[in5T]),
		args[5].(base.LazyFunc[in6T]),
		args[6].(base.LazyFunc[in7T]),
		args[7].(base.LazyFunc[in8T]),
	)
}
func (builder funcIn8[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn8[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn8 turns a function with eight input parameters into a lazy
// constructor.
func NewFuncIn8[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT any](
	src func(in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T) outT,
) Func {
	return Func{funcIn8[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT](func(
		lazyF1 base.LazyFunc[in1T], lazyF2 base.LazyFunc[in2T], lazyF3 base.LazyFunc[in3T], lazyF4 base.LazyFunc[in4T],
		lazyF5 base.LazyFunc[in5T], lazyF6 base.LazyFunc[in6T], lazyF7 base.LazyFunc[in7T], lazyF8 base.LazyFunc[in8T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6(), lazyF7(), lazyF8())
		}
	})}
}

// funcIn9 creates a lazy constructor based on a function with nine input
// parameters.
type funcIn9[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT any] func(
	base.LazyFunc[in1T],
	base.LazyFunc[in2T],
	base.LazyFunc[in3T],
	base.LazyFunc[in4T],
	base.LazyFunc[in5T],
	base.LazyFunc[in6T],
	base.LazyFunc[in7T],
	base.LazyFunc[in8T],
	base.LazyFunc[in9T],
) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with nine input parameters.
func (builder funcIn9[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(
		args[0].(base.LazyFunc[in1T]),
		args[1].(base.LazyFunc[in2T]),
		args[2].(base.LazyFunc[in3T]),
		args[3].(base.LazyFunc[in4T]),
		args[4].(base.LazyFunc[in5T]),
		args[5].(base.LazyFunc[in6T]),
		args[6].(base.LazyFunc[in7T]),
		args[7].(base.LazyFunc[in8T]),
		args[8].(base.LazyFunc[in9T]),
	)
}
func (builder funcIn9[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn9[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn9 turns a function with nine input parameters into a lazy
// constructor.
func NewFuncIn9[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT any](
	src func(in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T) outT,
) Func {
	return Func{funcIn9[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT](func(
		lazyF1 base.LazyFunc[in1T],
		lazyF2 base.LazyFunc[in2T],
		lazyF3 base.LazyFunc[in3T],
		lazyF4 base.LazyFunc[in4T],
		lazyF5 base.LazyFunc[in5T],
		lazyF6 base.LazyFunc[in6T],
		lazyF7 base.LazyFunc[in7T],
		lazyF8 base.LazyFunc[in8T],
		lazyF9 base.LazyFunc[in9T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6(), lazyF7(), lazyF8(), lazyF9())
		}
	})}
}

// funcIn10 creates a lazy constructor based on a function with ten input
// parameters.
type funcIn10[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT any,
] func(
	base.LazyFunc[in1T], base.LazyFunc[in2T], base.LazyFunc[in3T], base.LazyFunc[in4T],
	base.LazyFunc[in5T], base.LazyFunc[in6T], base.LazyFunc[in7T], base.LazyFunc[in8T],
	base.LazyFunc[in9T], base.LazyFunc[in10T],
) base.LazyFunc[outT]

// Build forms a lazy constructor for a function with ten input parameters.
func (builder funcIn10[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(
		args[0].(base.LazyFunc[in1T]),
		args[1].(base.LazyFunc[in2T]),
		args[2].(base.LazyFunc[in3T]),
		args[3].(base.LazyFunc[in4T]),
		args[4].(base.LazyFunc[in5T]),
		args[5].(base.LazyFunc[in6T]),
		args[6].(base.LazyFunc[in7T]),
		args[7].(base.LazyFunc[in8T]),
		args[8].(base.LazyFunc[in9T]),
		args[9].(base.LazyFunc[in10T]),
	)
}
func (builder funcIn10[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT]) GetOutType() reflect.Type {
	return reflect.TypeFor[outT]()
}
func (builder funcIn10[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT]) GetInputTypesHashSum() hashsum.Inputs {
	return hashsum.HashArgsByBuilder(builder)
}

// NewFuncIn10 turns a function with ten input parameters into a lazy
// constructor.
func NewFuncIn10[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT any](
	src func(in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T) outT,
) Func {
	return Func{funcIn10[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT](func(
		lazyF1 base.LazyFunc[in1T],
		lazyF2 base.LazyFunc[in2T],
		lazyF3 base.LazyFunc[in3T],
		lazyF4 base.LazyFunc[in4T],
		lazyF5 base.LazyFunc[in5T],
		lazyF6 base.LazyFunc[in6T],
		lazyF7 base.LazyFunc[in7T],
		lazyF8 base.LazyFunc[in8T],
		lazyF9 base.LazyFunc[in9T],
		lazyF10 base.LazyFunc[in10T],
	) base.LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6(), lazyF7(), lazyF8(), lazyF9(), lazyF10())
		}
	})}
}
