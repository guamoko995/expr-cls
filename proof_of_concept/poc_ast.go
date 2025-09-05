package poc

// Package poc demonstrates the speed difference between an expression compiled into
// a native Go function using closures versus one computed by the external library
// github.com/expr-lang/expr.

import (
	"reflect"
)

// Evaluation tree node represents a computation step in the tree.
type Node struct {
	// Operator is a closure representing some computation logic. It's hidden under
	// type any because the exact return type T is unknown at compile time but it's
	// always a zero-parameter function (func()T).
	operator any

	// Arguments are children nodes of this computational tree. Each argument can be
	// another evaluationTree or a data-provider operator whose Build method accepts
	// no arguments (like Const, Pointer, etc.). The count, order, and types of these
	// arguments must match those required by the operator's Build method.
	arguments []any
}

// Build constructs a function capable of evaluating the expression represented by
// this node.
func (et Node) Build() any {
	args := make([]reflect.Value, len(et.arguments))
	for i := 0; i < len(et.arguments); i++ {
		// Implicit recursion via reflection. Method(0) is always Build.
		args[i] = reflect.ValueOf(et.arguments[i]).
			Method(0).Call([]reflect.Value{})[0].Elem()
	}

	return reflect.ValueOf(et.operator).Method(0).Call(args)[0].Interface()
}

// BinaryOperator represents a binary operation used within the evaluation tree.
type BinaryOperator[arg1T, arg2T, resultT any] func(arg1T, arg2T) resultT

// Build creates a function that retrieves arguments through provided getter functions
// and applies the operator to them.
func (binary BinaryOperator[arg1T, arg2T, resultT]) Build(getArg1 func() arg1T,
	getArg2 func() arg2T) any {
	return func() resultT {
		return binary(getArg1(), getArg2())
	}
}

// UnaryOperator represents a unary operation used within the evaluation tree.
type UnaryOperator[argT, resultT any] func(argT) resultT

// Build creates a function that retrieves an argument through a provided getter
// function and applies the operator to it.
func (unary UnaryOperator[argT, resultT]) Build(getArg func() argT) any {
	return func() resultT {
		return unary(getArg())
	}
}

// Const holds a constant value for use in the evaluation tree.
type Const[T any] struct {
	Value T
}

// Build creates a function returning the stored constant value.
func (c Const[T]) Build() any {
	val := c.Value
	return func() T {
		return val
	}
}

// Pointer stores a pointer to a value for use in the evaluation tree.
type Pointer[T any] struct {
	ValuePtr *T
}

// Build creates a function that dereferences the pointer and returns its value.
func (p Pointer[T]) Build() any {
	return func() T {
		return *p.ValuePtr
	}
}

// Func wraps a plain Go function for use in the evaluation tree.
type Func[T any] func() T

// Build simply casts the receiver to any.
func (f Func[T]) Build() any {
	return f
}
