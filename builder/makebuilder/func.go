package makebuilder

// FuncIn1 transforms a function with one input parameter into a builder
// pattern. It takes a function that accepts an input of type in1T and returns
// output of type outT. The returned closure wraps the original function call
// using provided lazy evaluation functions.
func FuncIn1[in1T, outT any](source func(in1T) outT) func(func() in1T) func() outT {
	return func(f func() in1T) func() outT {
		return func() outT {
			return source(f())
		}
	}
}

// FuncIn2 transforms a function with two input parameters into a builder
// pattern. It takes a function that accepts inputs of types in1T and in2T and
// returns output of type outT. The returned closure wraps the original function
// call using provided lazy evaluation functions.
func FuncIn2[in1T, in2T, outT any](source func(in1T, in2T) outT) func(func() in1T, func() in2T) func() outT {
	return func(f1 func() in1T, f2 func() in2T) func() outT {
		return func() outT {
			return source(f1(), f2())
		}
	}
}

// FuncIn3 transforms a function with three input parameters into a builder
// pattern. It takes a function that accepts inputs of types in1T, in2T, and
// in3T and returns output of type outT. The returned closure wraps the original
// function call using provided lazy evaluation functions.
func FuncIn3[in1T, in2T, in3T, outT any](source func(in1T, in2T, in3T) outT) func(func() in1T, func() in2T, func() in3T) func() outT {
	return func(f1 func() in1T, f2 func() in2T, f3 func() in3T) func() outT {
		return func() outT {
			return source(f1(), f2(), f3())
		}
	}
}

// FuncIn4 transforms a function with four input parameters into a builder
// pattern. It takes a function that accepts inputs of types in1T, in2T, in3T,
// and in4T and returns output of type outT. The returned closure wraps the
// original function call using provided lazy evaluation functions.
func FuncIn4[in1T, in2T, in3T, in4T, outT any](source func(in1T, in2T, in3T, in4T) outT) func(func() in1T, func() in2T, func() in3T, func() in4T) func() outT {
	return func(f1 func() in1T, f2 func() in2T, f3 func() in3T, f4 func() in4T) func() outT {
		return func() outT {
			return source(f1(), f2(), f3(), f4())
		}
	}
}

// FuncIn5 transforms a function with five input parameters into a builder
// pattern. It takes a function that accepts inputs of types in1T, in2T, in3T,
// in4T, and in5T and returns output of type outT. The returned closure wraps
// the original function call using provided lazy evaluation functions.
func FuncIn5[in1T, in2T, in3T, in4T, in5T, outT any](source func(in1T, in2T, in3T, in4T, in5T) outT) func(func() in1T, func() in2T, func() in3T, func() in4T, func() in5T) func() outT {
	return func(f1 func() in1T, f2 func() in2T, f3 func() in3T, f4 func() in4T, f5 func() in5T) func() outT {
		return func() outT {
			return source(f1(), f2(), f3(), f4(), f5())
		}
	}
}

// FuncIn6 transforms a function with six input parameters into a builder
// pattern. It takes a function that accepts inputs of types in1T, in2T, ..., in6T
// and returns output of type outT. The returned closure wraps the original
// function call using provided lazy evaluation functions.
func FuncIn6[in1T, in2T, in3T, in4T, in5T, in6T, outT any](source func(in1T, in2T, in3T, in4T, in5T, in6T) outT) func(func() in1T, func() in2T, func() in3T, func() in4T, func() in5T, func() in6T) func() outT {
	return func(f1 func() in1T, f2 func() in2T, f3 func() in3T, f4 func() in4T, f5 func() in5T, f6 func() in6T) func() outT {
		return func() outT {
			return source(f1(), f2(), f3(), f4(), f5(), f6())
		}
	}
}

// FuncIn7 transforms a function with seven input parameters into a builder
// pattern. It takes a function that accepts inputs of types in1T, in2T, ..., in7T
// and returns output of type outT. The returned closure wraps the original
// function call using provided lazy evaluation functions.
func FuncIn7[in1T, in2T, in3T, in4T, in5T, in6T, in7T, outT any](source func(in1T, in2T, in3T, in4T, in5T, in6T, in7T) outT) func(func() in1T, func() in2T, func() in3T, func() in4T, func() in5T, func() in6T, func() in7T) func() outT {
	return func(f1 func() in1T, f2 func() in2T, f3 func() in3T, f4 func() in4T, f5 func() in5T, f6 func() in6T, f7 func() in7T) func() outT {
		return func() outT {
			return source(f1(), f2(), f3(), f4(), f5(), f6(), f7())
		}
	}
}

// FuncIn8 transforms a function with eight input parameters into a builder
// pattern. It takes a function that accepts inputs of types in1T, in2T, ..., in8T
// and returns output of type outT. The returned closure wraps the original
// function call using provided lazy evaluation functions.
func FuncIn8[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, outT any](source func(in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T) outT) func(func() in1T, func() in2T, func() in3T, func() in4T, func() in5T, func() in6T, func() in7T, func() in8T) func() outT {
	return func(f1 func() in1T, f2 func() in2T, f3 func() in3T, f4 func() in4T, f5 func() in5T, f6 func() in6T, f7 func() in7T, f8 func() in8T) func() outT {
		return func() outT {
			return source(f1(), f2(), f3(), f4(), f5(), f6(), f7(), f8())
		}
	}
}

// FuncIn9 transforms a function with nine input parameters into a builder
// pattern. It takes a function that accepts inputs of types in1T, in2T, ..., in9T
// and returns output of type outT. The returned closure wraps the original
// function call using provided lazy evaluation functions.
func FuncIn9[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, outT any](source func(in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T) outT) func(func() in1T, func() in2T, func() in3T, func() in4T, func() in5T, func() in6T, func() in7T, func() in8T, func() in9T) func() outT {
	return func(f1 func() in1T, f2 func() in2T, f3 func() in3T, f4 func() in4T, f5 func() in5T, f6 func() in6T, f7 func() in7T, f8 func() in8T, f9 func() in9T) func() outT {
		return func() outT {
			return source(f1(), f2(), f3(), f4(), f5(), f6(), f7(), f8(), f9())
		}
	}
}

// FuncIn10 transforms a function with ten input parameters into a builder
// pattern. It takes a function that accepts inputs of types in1T, in2T, ..., in10T
// and returns output of type outT. The returned closure wraps the original
// function call using provided lazy evaluation functions.
func FuncIn10[in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T, outT any](source func(in1T, in2T, in3T, in4T, in5T, in6T, in7T, in8T, in9T, in10T) outT) func(func() in1T, func() in2T, func() in3T, func() in4T, func() in5T, func() in6T, func() in7T, func() in8T, func() in9T, func() in10T) func() outT {
	return func(f1 func() in1T, f2 func() in2T, f3 func() in3T, f4 func() in4T, f5 func() in5T, f6 func() in6T, f7 func() in7T, f8 func() in8T, f9 func() in9T, f10 func() in10T) func() outT {
		return func() outT {
			return source(f1(), f2(), f3(), f4(), f5(), f6(), f7(), f8(), f9(), f10())
		}
	}
}
