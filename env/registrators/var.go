package registrators

import (
	"reflect"
	"unsafe"

	base "github.com/guamoko995/expr-cls/env/base"
	"github.com/guamoko995/expr-cls/internal/hashsum"
)

type VariableType struct {
	builderMaker GenericBuildersMaker
}

func (v VariableType) MakeGetVarBuildersFromSrc(srcTR reflect.Type) map[string]base.Builder {
	return v.builderMaker.MakeGetVarBuildersFromSrc(srcTR)
}

func (v VariableType) MakeConstBuilder(val any) base.Builder {
	return v.builderMaker.MakeConstBuilder(val)
}

func (v VariableType) GetVarType() reflect.Type {
	return v.builderMaker.GetVarType()
}

// NewVar turns a value into a lazy constructor.
func NewVar[T any]() VariableType {
	return VariableType{builderMaker[T]{}}
}

type GenericBuildersMaker interface {
	MakeGetVarBuildersFromSrc(srcTR reflect.Type) map[string]base.Builder
	MakeConstBuilder(val any) base.Builder
	GetVarType() reflect.Type
}

type builderMaker[varT any] struct{}

func (builderMaker[varT]) MakeGetVarBuildersFromSrc(srcTR reflect.Type) map[string]base.Builder {
	results := make(map[string]base.Builder, srcTR.NumField())
	for i := range srcTR.NumField() {
		if srcTR.Field(i).Type == reflect.TypeFor[varT]() {
			name := srcTR.Field(i).Name
			offset := srcTR.Field(i).Offset
			results[name] = makeGetVarBuilder[varT](offset)
		}
	}
	return results
}

func makeGetVarBuilder[varT any](offset uintptr) varBuilder[varT] {
	return func(pSrc unsafe.Pointer) base.LazyFunc[varT] {
		varPtr := (*varT)(unsafe.Pointer(uintptr(pSrc) + offset))
		return func() varT { return *varPtr }
	}
}

type varBuilder[varT any] func(pSrc unsafe.Pointer) base.LazyFunc[varT]

func (builder varBuilder[varT]) Build(args []base.GenericLazyFunc) base.GenericLazyFunc {
	return builder(args[0].(unsafe.Pointer))
}

func (builder varBuilder[varT]) GetOutType() reflect.Type {
	return reflect.TypeFor[varT]()
}
func (builder varBuilder[varT]) GetInputTypesHashSum() hashsum.Inputs {
	return 0
}

func (builderMaker[varT]) MakeConstBuilder(val any) base.Builder {
	v := val.(varT)
	return NewConst(v)
}

func (builderMaker[varT]) GetVarType() reflect.Type {
	return reflect.TypeFor[varT]()
}

type VarSourse struct {
	src Sourse
}

func (src VarSourse) GetSrcType() reflect.Type {
	return src.src.GetSrcType()
}

type Sourse interface {
	GetSrcType() reflect.Type
}

type sourse[srcT any] struct{}

func (sourse[srcT]) GetSrcType() reflect.Type {
	return reflect.TypeFor[srcT]()
}

// NewVar turns a value into a lazy constructor.
func NewVarSourse[T any]() VarSourse {
	return VarSourse{sourse[T]{}}
}
