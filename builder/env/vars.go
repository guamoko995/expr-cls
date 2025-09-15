package env

import (
	"maps"
	"reflect"
	"unsafe"

	"github.com/guamoko995/expr-cls/builder/base"
)

func (env *Env) RegDefaultVarTypes() {
	RegVarType[int](env)
	RegVarType[float64](env)
	RegVarType[string](env)
	RegVarType[bool](env)
}

func RegVarType[varT any](env *Env) {
	env.BuilderMakers[reflect.TypeFor[varT]()] = builderMaker[varT]{}

	for srcTR := range env.VarBuilders {
		maps.Copy(env.VarBuilders[srcTR], (builderMaker[varT]{}).MakeGetVarBuildersFromSrc(srcTR))
	}
}

func RegVarSrc[srcT any](env ...*Env) {
	if len(env) == 0 {
		env = append(env, Global)
	}
	for i := range env {
		srcTR := reflect.TypeFor[srcT]()
		env[i].VarBuilders[srcTR] = make(map[string]base.Builder)
		for _, varBuildersMaker := range env[i].BuilderMakers {
			maps.Copy(env[i].VarBuilders[srcTR], varBuildersMaker.MakeGetVarBuildersFromSrc(srcTR))
		}
	}
}

type getVarBuilder[varT any] func(pSrc unsafe.Pointer) base.LazyFunc[varT]

func (builder getVarBuilder[varT]) Build(args []any) base.GenericLazyFunc {
	return builder(args[0].(unsafe.Pointer))
}

type getConstBuilder[varT any] func() base.LazyFunc[varT]

func (builder getConstBuilder[varT]) Build(args []any) base.GenericLazyFunc {
	return builder()
}

func makeGetVarBuilder[varT any](offset uintptr) getVarBuilder[varT] {
	return func(pSrc unsafe.Pointer) base.LazyFunc[varT] {
		varPtr := (*varT)(unsafe.Pointer(uintptr(pSrc) + offset))
		return func() varT { return *varPtr }
	}
}

type BuildersMaker interface {
	MakeGetVarBuildersFromSrc(srcTR reflect.Type) map[string]base.Builder
	MakeConstBuilder(val any) base.Builder
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

func (builderMaker[varT]) MakeConstBuilder(val any) base.Builder {
	v := val.(varT)
	return getConstBuilder[varT](func() base.LazyFunc[varT] {
		return func() varT {
			return v
		}
	})
}
