package base

import (
	"reflect"

	"github.com/guamoko995/expr-cls/internal/hashsum"
)

// LazyFunc represents a lazy-evaluated function.
type LazyFunc[T any] func() T

// Universal type alias for representing any lazy function.
type GenericLazyFunc any

type Builder interface {
	// Build creates a new instance of a lazy function.
	Build([]GenericLazyFunc) GenericLazyFunc

	// GetOutType возвращает отражение типа возвращаемого значения ленивой
	// функции.
	GetOutType() reflect.Type

	// GetInputTypesHashSum возвращает хэш-сумму типов возвращаемых ленивыми
	// функциями, являющимися аргументами билдера.
	GetInputTypesHashSum() hashsum.Inputs
}
