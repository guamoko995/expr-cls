package env

import (
	"maps"
	"reflect"
	"unsafe"

	"github.com/guamoko995/expr-cls/builder/base"
)

func (env *Env) SetDefaultVarTypes() {
	DefVarType[int](env)
	DefVarType[float64](env)
	DefVarType[string](env)
	DefVarType[bool](env)
}

func DefVarType[varT any](env *Env) {
	env.VarBuilderMakers[reflect.TypeFor[varT]()] = getVarBuildersMaker[varT]{}

	for srcTR := range env.VarBuilders {
		maps.Copy(env.VarBuilders[srcTR], (getVarBuildersMaker[varT]{}).Make(srcTR))
	}
}

func DefVarSrc[srcT any](env *Env) {
	srcTR := reflect.TypeFor[srcT]()
	env.VarBuilders[srcTR] = make(map[string]base.Builder)
	for _, varBuildersMaker := range env.VarBuilderMakers {
		maps.Copy(env.VarBuilders[srcTR], varBuildersMaker.Make(srcTR))
	}

}

type getVarBuilder[varT any] func(pSrc unsafe.Pointer) base.LazyFunc[varT]

func (builder getVarBuilder[varT]) Build(args []any) base.GenericLazyFunc {
	return builder(args[0].(unsafe.Pointer))
}

func makeGetVarBuilder[varT any](offset uintptr) getVarBuilder[varT] {
	return func(pSrc unsafe.Pointer) base.LazyFunc[varT] {
		varPtr := (*varT)(unsafe.Pointer(uintptr(pSrc) + offset))
		return func() varT { return *varPtr }
	}
}

type anyGetVarBuildersMaker interface {
	Make(srcTR reflect.Type) map[string]base.Builder
}

type getVarBuildersMaker[varT any] struct{}

func (getVarBuildersMaker[varT]) Make(srcTR reflect.Type) map[string]base.Builder {
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
