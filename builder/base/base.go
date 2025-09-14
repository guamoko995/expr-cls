package base

// LazyFunc represents a lazy-evaluated function.
type LazyFunc[T any] func() T

// Universal type alias for representing any lazy function.
type GenericLazyFunc any

type Builder interface {
	// Build creates a new instance of a lazy function.
	Build([]any) GenericLazyFunc
}
