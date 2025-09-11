package makebuilder

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/guamoko995/expr-cls/internal/hash"
)

var ExpectedVarTypes = map[hash.Type]any{
	hash.HashType[int]():    VarBuilderMaker[int]{},
	hash.HashType[string](): VarBuilderMaker[string]{},
	hash.HashType[bool]():   VarBuilderMaker[bool]{},
}

func Sources[srcT any]() func(pSource unsafe.Pointer) map[string]any {
	sources := make(map[hash.Type]map[string]any)
	builders := make([]func(unsafe.Pointer) map[string]any, 0, len(ExpectedVarTypes))
	for varTHash, varBuilderMaker := range ExpectedVarTypes {
		if _, exist := sources[varTHash]; !exist {
			sources[varTHash] = make(map[string]any)
		}
		builder := reflect.ValueOf(varBuilderMaker).Method(0).Call([]reflect.Value{reflect.ValueOf(reflect.TypeFor[srcT]())})[0].Interface().(func(unsafe.Pointer) map[string]any)

		builders = append(builders, builder)
	}
	return func(pSource unsafe.Pointer) map[string]any {
		m := make(map[string]any)

		for i := range builders {
			for varName, fn := range builders[i](pSource) {
				m[varName] = fn
			}
		}
		return m
	}

}

type VarBuilderMaker[varT any] struct{}

// Vars принимает отражение типа структуры-источника (значения полей
// структуры-источника в конечном счете и будут значениями переменных в
// выражении).
// Vars возвращает некоторую функцию Builder.
//
// Builder принимает опасный указатель на структуру-источник.
// Builder возвращает словарь замыканий.
//
// В словаре замыканий, ключом является семантика доступа к полю
// структуры-источника; а значением является замыкание.
//
// Конечное замыкание, не принимает аргументы и возвращает значение типа varT,
// из соответствующего поля структуры-источника.
func (VarBuilderMaker[varT]) Vars(strTR reflect.Type) func(unsafe.Pointer) map[string]any {
	offsets := make(map[string]uintptr)
	for i := range strTR.NumField() {
		if strTR.Field(i).Type == reflect.TypeFor[varT]() {
			name := strTR.Field(i).Name
			offset := strTR.Field(i).Offset
			offsets[name] = offset
		}
	}
	return func(ptr unsafe.Pointer) map[string]any {
		result := make(map[string]any, len(offsets))
		for fieldName, offset := range offsets {
			result[fieldName] = func() varT {
				return *(*varT)(unsafe.Pointer(uintptr(ptr) + offset))
			}
		}
		return result
	}
}

type argsExample struct {
	A int
	a int
	B bool
	b bool
	C string
}

// TODO benchmark
func Test(t *testing.T) {
	srces := Sources[argsExample]()

	v := argsExample{A: 1, a: 20, B: true, b: false, C: "rrr"}

	fns := srces(unsafe.Pointer(&v))

	for fieldName, fn := range fns {
		var val any
		switch fn := fn.(type) {
		case func() int:
			val = fn()
		case func() bool:
			val = fn()
		case func() string:
			val = fn()
		}

		fmt.Printf("%s: %v\n", fieldName, val)
	}
}
