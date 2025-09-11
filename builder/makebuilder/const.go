package makebuilder

func Const[T any](val T) func() func() T {
	return func() func() T {
		return func() T {
			return val
		}
	}
}
