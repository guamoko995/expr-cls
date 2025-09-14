package base

// FuncIn1Builder creates a lazy constructor based on a function with one input
// parameter.
type FuncIn1Builder[in1T, outT any] func(LazyFunc[in1T]) LazyFunc[outT]

// Build forms a lazy constructor for a function with one input parameter.
func (builder FuncIn1Builder[in1T, outT]) Build(args []any) GenericLazyFunc {
	return builder(args[0].(LazyFunc[in1T]))
}

// MakeFuncIn1Builder turns a function with one input parameter into a lazy
// constructor.
func MakeFuncIn1Builder[in1T, outT any](src func(in1T) outT) FuncIn1Builder[in1T, outT] {
	return func(lazyF LazyFunc[in1T]) LazyFunc[outT] {
		return func() outT {
			return src(lazyF())
		}
	}
}

// FuncIn2Builder creates a lazy constructor based on a function with two input
// parameters.
type FuncIn2Builder[in1T, in2T, outT any] func(LazyFunc[in1T], LazyFunc[in2T]) LazyFunc[outT]

// Build forms a lazy constructor for a function with two input parameters.
func (builder FuncIn2Builder[in1T, in2T, outT]) Build(args []any) GenericLazyFunc {
	return builder(args[0].(LazyFunc[in1T]), args[1].(LazyFunc[in2T]))
}

// MakeFuncIn2Builder turns a function with two input parameters into a lazy
// constructor.
func MakeFuncIn2Builder[in1T, in2T, outT any](src func(in1T, in2T) outT) FuncIn2Builder[in1T, in2T, outT] {
	return func(lazyF1 LazyFunc[in1T], lazyF2 LazyFunc[in2T]) LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2())
		}
	}
}

// FuncIn3Builder creates a lazy constructor based on a function with three input
// parameters.
type FuncIn3Builder[in1T, in2T, in3T, outT any] func(LazyFunc[in1T], LazyFunc[in2T], LazyFunc[in3T]) LazyFunc[outT]

// Build forms a lazy constructor for a function with three input parameters.
func (builder FuncIn3Builder[in1T, in2T, in3T, outT]) Build(args []any) GenericLazyFunc {
	return builder(args[0].(LazyFunc[in1T]), args[1].(LazyFunc[in2T]), args[2].(LazyFunc[in3T]))
}

// MakeFuncIn3Builder turns a function with three input parameters into a lazy
// constructor.
func MakeFuncIn3Builder[in1T, in2T, in3T, outT any](src func(in1T, in2T, in3T) outT) FuncIn3Builder[in1T, in2T, in3T, outT] {
	return func(lazyF1 LazyFunc[in1T], lazyF2 LazyFunc[in2T], lazyF3 LazyFunc[in3T]) LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3())
		}
	}
}

// FuncIn4Builder creates a lazy constructor based on a function with four input
// parameters.
type FuncIn4Builder[in1T, in2T, in3T, in4T, outT any] func(LazyFunc[in1T], LazyFunc[in2T], LazyFunc[in3T], LazyFunc[in4T]) LazyFunc[outT]

// Build forms a lazy constructor for a function with four input parameters.
func (builder FuncIn4Builder[in1T, in2T, in3T, in4T, outT]) Build(args []any) GenericLazyFunc {
	return builder(args[0].(LazyFunc[in1T]), args[1].(LazyFunc[in2T]), args[2].(LazyFunc[in3T]), args[3].(LazyFunc[in4T]))
}

// MakeFuncIn4Builder turns a function with four input parameters into a lazy
// constructor.
func MakeFuncIn4Builder[in1T, in2T, in3T, in4T, outT any](src func(in1T, in2T, in3T, in4T) outT) FuncIn4Builder[in1T, in2T, in3T, in4T, outT] {
	return func(lazyF1 LazyFunc[in1T], lazyF2 LazyFunc[in2T], lazyF3 LazyFunc[in3T], lazyF4 LazyFunc[in4T]) LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4())
		}
	}
}

// FuncIn5Builder creates a lazy constructor based on a function with five input
// parameters.
type FuncIn5Builder[in1T, in2T, in3T, in4T, in5T, outT any] func(LazyFunc[in1T], LazyFunc[in2T], LazyFunc[in3T], LazyFunc[in4T], LazyFunc[in5T]) LazyFunc[outT]

// Build forms a lazy constructor for a function with five input parameters.
func (builder FuncIn5Builder[in1T, in2T, in3T, in4T, in5T, outT]) Build(args []any) GenericLazyFunc {
	return builder(args[0].(LazyFunc[in1T]), args[1].(LazyFunc[in2T]), args[2].(LazyFunc[in3T]), args[3].(LazyFunc[in4T]), args[4].(LazyFunc[in5T]))
}

// MakeFuncIn5Builder turns a function with five input parameters into a lazy
// constructor.
func MakeFuncIn5Builder[in1T, in2T, in3T, in4T, in5T, outT any](src func(in1T, in2T, in3T, in4T, in5T) outT) FuncIn5Builder[in1T, in2T, in3T, in4T, in5T, outT] {
	return func(lazyF1 LazyFunc[in1T], lazyF2 LazyFunc[in2T], lazyF3 LazyFunc[in3T], lazyF4 LazyFunc[in4T], lazyF5 LazyFunc[in5T]) LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5())
		}
	}
}

// FuncIn6Builder creates a lazy constructor based on a function with six input
// parameters.
type FuncIn6Builder[in1T, in2T, in3T, in4T, in5T, in6T, outT any] func(LazyFunc[in1T], LazyFunc[in2T], LazyFunc[in3T], LazyFunc[in4T], LazyFunc[in5T], LazyFunc[in6T]) LazyFunc[outT]

// Build forms a lazy constructor for a function with six input parameters.
func (builder FuncIn6Builder[in1T, in2T, in3T, in4T, in5T, in6T, outT]) Build(args []any) GenericLazyFunc {
	return builder(args[0].(LazyFunc[in1T]), args[1].(LazyFunc[in2T]), args[2].(LazyFunc[in3T]), args[3].(LazyFunc[in4T]), args[4].(LazyFunc[in5T]), args[5].(LazyFunc[in6T]))
}

// MakeFuncIn6Builder turns a function with six input parameters into a lazy
// constructor.
func MakeFuncIn6Builder[in1T, in2T, in3T, in4T, in5T, in6T, outT any](src func(in1T, in2T, in3T, in4T, in5T, in6T) outT) FuncIn6Builder[in1T, in2T, in3T, in4T, in5T, in6T, outT] {
	return func(lazyF1 LazyFunc[in1T], lazyF2 LazyFunc[in2T], lazyF3 LazyFunc[in3T], lazyF4 LazyFunc[in4T], lazyF5 LazyFunc[in5T], lazyF6 LazyFunc[in6T]) LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6())
		}
	}
}

// FuncIn7Builder creates a lazy constructor based on a function with seven input
// parameters.
type FuncIn7Builder[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT any] func(LazyFunc[in1T], LazyFunc[in2T], LazyFunc[in3T], LazyFunc[in4T], LazyFunc[in5T], LazyFunc[in6T], LazyFunc[in7T]) LazyFunc[outT]

// Build forms a lazy constructor for a function with seven input parameters.
func (builder FuncIn7Builder[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT]) Build(args []any) GenericLazyFunc {
	return builder(args[0].(LazyFunc[in1T]), args[1].(LazyFunc[in2T]), args[2].(LazyFunc[in3T]), args[3].(LazyFunc[in4T]), args[4].(LazyFunc[in5T]), args[5].(LazyFunc[in6T]), args[6].(LazyFunc[in7T]))
}

// MakeFuncIn7Builder turns a function with seven input parameters into a lazy
// constructor.
func MakeFuncIn7Builder[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT any](src func(in1T, in2T, in3T, in4T, in5T, in6T, in7T) outT) FuncIn7Builder[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT] {
	return func(lazyF1 LazyFunc[in1T], lazyF2 LazyFunc[in2T], lazyF3 LazyFunc[in3T], lazyF4 LazyFunc[in4T], lazyF5 LazyFunc[in5T], lazyF6 LazyFunc[in6T], lazyF7 LazyFunc[in7T]) LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6(), lazyF7())
		}
	}
}

// FuncIn8Builder creates a lazy constructor based on a function with eight input
// parameters.
type FuncIn8Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT any,
] func(
	LazyFunc[in1T], LazyFunc[in2T], LazyFunc[in3T], LazyFunc[in4T],
	LazyFunc[in5T], LazyFunc[in6T], LazyFunc[in7T], LazyFunc[in8T],
) LazyFunc[outT]

// Build forms a lazy constructor for a function with eight input parameters.
func (builder FuncIn8Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT,
]) Build(args []any) GenericLazyFunc {
	return builder(
		args[0].(LazyFunc[in1T]), args[1].(LazyFunc[in2T]), args[2].(LazyFunc[in3T]), args[3].(LazyFunc[in4T]),
		args[4].(LazyFunc[in5T]), args[5].(LazyFunc[in6T]), args[6].(LazyFunc[in7T]), args[7].(LazyFunc[in8T]),
	)
}

// MakeFuncIn8Builder turns a function with eight input parameters into a lazy
// constructor.
func MakeFuncIn8Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT any,
](src func(in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T) outT) FuncIn8Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT,
] {
	return func(
		lazyF1 LazyFunc[in1T], lazyF2 LazyFunc[in2T], lazyF3 LazyFunc[in3T], lazyF4 LazyFunc[in4T],
		lazyF5 LazyFunc[in5T], lazyF6 LazyFunc[in6T], lazyF7 LazyFunc[in7T], lazyF8 LazyFunc[in8T],
	) LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6(), lazyF7(), lazyF8())
		}
	}
}

// FuncIn9Builder creates a lazy constructor based on a function with nine input
// parameters.
type FuncIn9Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT any,
] func(
	LazyFunc[in1T], LazyFunc[in2T], LazyFunc[in3T], LazyFunc[in4T],
	LazyFunc[in5T], LazyFunc[in6T], LazyFunc[in7T], LazyFunc[in8T],
	LazyFunc[in9T],
) LazyFunc[outT]

// Build forms a lazy constructor for a function with nine input parameters.
func (builder FuncIn9Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT,
]) Build(args []any) GenericLazyFunc {
	return builder(
		args[0].(LazyFunc[in1T]), args[1].(LazyFunc[in2T]), args[2].(LazyFunc[in3T]), args[3].(LazyFunc[in4T]),
		args[4].(LazyFunc[in5T]), args[5].(LazyFunc[in6T]), args[6].(LazyFunc[in7T]), args[7].(LazyFunc[in8T]),
		args[8].(LazyFunc[in9T]),
	)
}

// MakeFuncIn9Builder turns a function with nine input parameters into a lazy
// constructor.
func MakeFuncIn9Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT any,
](src func(in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T) outT) FuncIn9Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT,
] {
	return func(
		lazyF1 LazyFunc[in1T], lazyF2 LazyFunc[in2T], lazyF3 LazyFunc[in3T], lazyF4 LazyFunc[in4T],
		lazyF5 LazyFunc[in5T], lazyF6 LazyFunc[in6T], lazyF7 LazyFunc[in7T], lazyF8 LazyFunc[in8T],
		lazyF9 LazyFunc[in9T],
	) LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6(), lazyF7(), lazyF8(), lazyF9())
		}
	}
}

// FuncIn10Builder creates a lazy constructor based on a function with ten input
// parameters.
type FuncIn10Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT any,
] func(
	LazyFunc[in1T], LazyFunc[in2T], LazyFunc[in3T], LazyFunc[in4T],
	LazyFunc[in5T], LazyFunc[in6T], LazyFunc[in7T], LazyFunc[in8T],
	LazyFunc[in9T], LazyFunc[in10T],
) LazyFunc[outT]

// Build forms a lazy constructor for a function with ten input parameters.
func (builder FuncIn10Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT,
]) Build(args []any) GenericLazyFunc {
	return builder(
		args[0].(LazyFunc[in1T]), args[1].(LazyFunc[in2T]), args[2].(LazyFunc[in3T]), args[3].(LazyFunc[in4T]),
		args[4].(LazyFunc[in5T]), args[5].(LazyFunc[in6T]), args[6].(LazyFunc[in7T]), args[7].(LazyFunc[in8T]),
		args[8].(LazyFunc[in9T]), args[9].(LazyFunc[in10T]),
	)
}

// MakeFuncIn10Builder turns a function with ten input parameters into a lazy
// constructor.
func MakeFuncIn10Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT any,
](src func(in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T) outT) FuncIn10Builder[
	in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT,
] {
	return func(
		lazyF1 LazyFunc[in1T], lazyF2 LazyFunc[in2T], lazyF3 LazyFunc[in3T], lazyF4 LazyFunc[in4T],
		lazyF5 LazyFunc[in5T], lazyF6 LazyFunc[in6T], lazyF7 LazyFunc[in7T], lazyF8 LazyFunc[in8T],
		lazyF9 LazyFunc[in9T], lazyF10 LazyFunc[in10T],
	) LazyFunc[outT] {
		return func() outT {
			return src(lazyF1(), lazyF2(), lazyF3(), lazyF4(), lazyF5(), lazyF6(), lazyF7(), lazyF8(), lazyF9(), lazyF10())
		}
	}
}
