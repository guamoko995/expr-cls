package ast

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/guamoko995/expr-cls/env"
	basepkg "github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/file"
	"github.com/guamoko995/expr-cls/internal/hashsum"
)

// Node represents items of abstract syntax tree.
type Node interface {
	Location() file.Location
	SetLocation(file.Location)
	String() string
	IsConstant() bool
	Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error)
}

// Patch replaces the node with a new one.
// Location information is preserved.
// Type information is lost.
func Patch(node *Node, newNode Node) {
	newNode.SetLocation((*node).Location())
	*node = newNode
}

// base is a base struct for all nodes.
type base struct {
	loc     file.Location
	isConst bool
}

func (n *base) IsConstant() bool {
	return n.isConst
}

// Location returns the location of the node in the source code.
func (n *base) Location() file.Location {
	return n.loc
}

// SetLocation sets the location of the node in the source code.
func (n *base) SetLocation(loc file.Location) {
	n.loc = loc
}

func (n *base) Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error) {
	return nil, errors.New("not emplemented")
}

// NilNode represents nil.
type NilNode struct {
	base
}

// IdentifierNode represents an identifier.
type IdentifierNode struct {
	base
	Value string // Name of the identifier. Like "foo" in "foo.bar".
}

func (n *IdentifierNode) Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error) {
	srcTR := reflect.TypeOf(varSrc).Elem()

	if builder, exist := env.Variables[srcTR][n.Value]; exist {
		return builder.Build([]basepkg.GenericLazyFunc{unsafe.Pointer(reflect.ValueOf(varSrc).Pointer())}), nil
	}

	if builder, exist := env.Const[n.Value]; exist {
		n.isConst = true
		return builder.Build([]basepkg.GenericLazyFunc{}), nil
	}

	return nil, fmt.Errorf("identifier %q not found", n.Value)
}

// IntegerNode represents an integer.
type IntegerNode struct {
	base
	Value int // Value of the integer.
}

func (n *IntegerNode) Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error) {
	n.isConst = true
	return basepkg.LazyFunc[int](func() int { return n.Value }), nil
}

// FloatNode represents a float.
type FloatNode struct {
	base
	Value float64 // Value of the float.
}

func (n *FloatNode) Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error) {
	n.isConst = true
	return basepkg.LazyFunc[float64](func() float64 { return n.Value }), nil
}

// BoolNode represents a boolean.
type BoolNode struct {
	base
	Value bool // Value of the boolean.
}

func (n *BoolNode) Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error) {
	n.isConst = true
	return basepkg.LazyFunc[bool](func() bool { return n.Value }), nil
}

// StringNode represents a string.
type StringNode struct {
	base
	Value string // Value of the string.
}

func (n *StringNode) Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error) {
	n.isConst = true
	return basepkg.LazyFunc[string](func() string { return n.Value }), nil
}

// ConstantNode represents a constant.
// Constants are predefined values like nil, true, false, array, map, etc.
// The parser.Parse will never generate ConstantNode, it is only generated
// by the optimizer.
type ConstantNode struct {
	base
	Value any // Value of the constant.
}

// UnaryNode represents a unary operator.
type UnaryNode struct {
	base
	Operator string // Operator of the unary operator. Like "!" in "!foo" or "not" in "not foo".
	Node     Node   // Node of the unary operator. Like "foo" in "!foo".
}

func (n *UnaryNode) Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error) {
	arg1, err := n.Node.Build(env, varSrc)
	if err != nil {
		return nil, err
	}
	n.isConst = n.Node.IsConstant()

	if _, exist := env.Unary[n.Operator]; !exist {
		return nil, fmt.Errorf("environment does not contain any implementations of the %q operator", n.Operator)
	}

	argsHash := hashsum.HashArgs(reflect.TypeOf(arg1).Out(0))

	fn, exist := env.Unary[n.Operator][argsHash]
	if !exist {
		return nil, fmt.Errorf("environment does not contain implementations of the %q operator for the given arguments", n.Operator)
	}

	result := fn.Build([]basepkg.GenericLazyFunc{arg1})

	if n.IsConstant() {
		result := reflect.ValueOf(result).Call([]reflect.Value{})[0].Interface()
		resultTR := reflect.TypeOf(result)

		if builderMaker, exist := env.VariableMakers[resultTR]; exist {
			return builderMaker.MakeConstBuilder(result).Build(nil), nil
		}
	}

	return result, nil
}

// BinaryNode represents a binary operator.
type BinaryNode struct {
	base
	Operator string // Operator of the binary operator. Like "+" in "foo + bar" or "matches" in "foo matches bar".
	Left     Node   // Left node of the binary operator.
	Right    Node   // Right node of the binary operator.
}

func (n *BinaryNode) Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error) {

	arg1, err := n.Left.Build(env, varSrc)
	if err != nil {
		return nil, err
	}

	arg2, err := n.Right.Build(env, varSrc)
	if err != nil {
		return nil, err
	}

	n.isConst = n.Right.IsConstant() && n.Left.IsConstant()

	if _, exist := env.Binary[n.Operator]; !exist {
		return nil, fmt.Errorf("environment does not contain any implementations of the %q operator", n.Operator)
	}

	argsHash := hashsum.HashArgs(reflect.TypeOf(arg1).Out(0), reflect.TypeOf(arg2).Out(0))

	fn, exist := env.Binary[n.Operator][argsHash]
	if !exist {
		return nil, fmt.Errorf("environment does not contain implementations of the %q operator for the given arguments", n.Operator)
	}

	result := fn.Build([]basepkg.GenericLazyFunc{arg1, arg2})

	if n.IsConstant() {
		result := reflect.ValueOf(result).Call([]reflect.Value{})[0].Interface()
		resultTR := reflect.TypeOf(result)

		if builderMaker, exist := env.VariableMakers[resultTR]; exist {
			return builderMaker.MakeConstBuilder(result).Build(nil), nil
		}
	}

	return result, nil
}

// ChainNode represents an optional chaining group.
// A few MemberNode nodes can be chained together,
// and will be wrapped in a ChainNode. Example:
//
//	foo.bar?.baz?.qux
//
// The whole chain will be wrapped in a ChainNode.
type ChainNode struct {
	base
	Node Node // Node of the chain.
}

// MemberNode represents a member access.
// It can be a field access, a method call,
// or an array element access.
// Example:
//
//	foo.bar or foo["bar"]
//	foo.bar()
//	array[0]
type MemberNode struct {
	base
	Node     Node // Node of the member access. Like "foo" in "foo.bar".
	Property Node // Property of the member access. For property access it is a StringNode.
	Optional bool // If true then the member access is optional. Like "foo?.bar".
	Method   bool
}

// SliceNode represents access to a slice of an array.
// Example:
//
//	array[1:4]
type SliceNode struct {
	base
	Node Node // Node of the slice. Like "array" in "array[1:4]".
	From Node // From an index of the array. Like "1" in "array[1:4]".
	To   Node // To an index of the array. Like "4" in "array[1:4]".
}

// CallNode represents a function or a method call.
type CallNode struct {
	base
	Callee    Node   // Node of the call. Like "foo" in "foo()".
	Arguments []Node // Arguments of the call.
}

func (n *CallNode) Build(env *env.Enviroment, varSrc any) (basepkg.GenericLazyFunc, error) {
	ident, ok := n.Callee.(*IdentifierNode)
	if !ok {
		return nil, errors.New("unexpected call")
	}

	_, exist := env.Func[ident.Value]
	if !exist {
		return nil, fmt.Errorf("function %q not found", ident.Value)
	}

	args := make([]basepkg.GenericLazyFunc, len(n.Arguments))
	argTypes := make([]reflect.Type, len(n.Arguments))
	n.isConst = true
	for i := range len(n.Arguments) {
		var err error
		args[i], err = n.Arguments[i].Build(env, varSrc)
		if err != nil {
			return nil, err
		}
		if !n.Arguments[i].IsConstant() {
			n.isConst = false
		}
		argTypes[i] = reflect.TypeOf(args[i]).Out(0)
	}

	fn, exist := env.Func[ident.Value][hashsum.HashArgs(argTypes...)]
	if !exist {
		return nil, fmt.Errorf("environment does not contain implementations of the %q function for the given arguments", ident.Value)
	}

	result := fn.Build(args)

	if n.IsConstant() {
		result := reflect.ValueOf(result).Call([]reflect.Value{})[0].Interface()
		resultTR := reflect.TypeOf(result)

		if builderMaker, exist := env.VariableMakers[resultTR]; exist {
			return builderMaker.MakeConstBuilder(result).Build(nil), nil
		}
	}

	return result, nil
}

/*/ BuiltinNode represents a builtin function call.
type BuiltinNode struct {
	base
	Name      string // Name of the builtin function. Like "len" in "len(foo)".
	Arguments []Node // Arguments of the builtin function.
}//*/

// PredicateNode represents a predicate.
// Example:
//
//	filter(foo, .bar == 1)
//
// The predicate is ".bar == 1".
type PredicateNode struct {
	base
	Node Node // Node of the predicate body.
}

// PointerNode represents a pointer to a current value in predicate.
type PointerNode struct {
	base
	Name string // Name of the pointer. Like "index" in "#index".
}

// ConditionalNode represents a ternary operator.
type ConditionalNode struct {
	base
	Cond Node // Condition of the ternary operator. Like "foo" in "foo ? bar : baz".
	Exp1 Node // Expression 1 of the ternary operator. Like "bar" in "foo ? bar : baz".
	Exp2 Node // Expression 2 of the ternary operator. Like "baz" in "foo ? bar : baz".
}

// VariableDeclaratorNode represents a variable declaration.
type VariableDeclaratorNode struct {
	base
	Name  string // Name of the variable. Like "foo" in "let foo = 1; foo + 1".
	Value Node   // Value of the variable. Like "1" in "let foo = 1; foo + 1".
	Expr  Node   // Expression of the variable. Like "foo + 1" in "let foo = 1; foo + 1".
}

// SequenceNode represents a sequence of nodes separated by semicolons.
// All nodes are executed, only the last node will be returned.
type SequenceNode struct {
	base
	Nodes []Node
}

// ArrayNode represents an array.
type ArrayNode struct {
	base
	Nodes []Node // Nodes of the array.
}

// MapNode represents a map.
type MapNode struct {
	base
	Pairs []Node // PairNode nodes.
}

// PairNode represents a key-value pair of a map.
type PairNode struct {
	base
	Key   Node // Key of the pair.
	Value Node // Value of the pair.
}
