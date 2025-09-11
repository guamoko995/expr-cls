package hash

import (
	"hash/fnv"
	"reflect"
)

type Args uint64
type Type uint64

func HashArgsByFunc(fn any) Args {
	h := fnv.New64()
	fnT := reflect.TypeOf(fn)
	if fnT.Kind() != reflect.Func {
		panic("unexpected argument")
	}
	for i := range fnT.NumIn() {
		_, err := h.Write([]byte(fnT.In(i).String()))
		if err != nil {
			panic(err)
		}
	}
	return Args(h.Sum64())
}

func HashArgsByBuilder(fn any) Args {
	h := fnv.New64()
	fnT := reflect.TypeOf(fn)
	if fnT.Kind() != reflect.Func {
		panic("unexpected argument")
	}
	for i := range fnT.NumIn() {
		if fnT.In(i).Kind() != reflect.Func || fnT.In(i).NumOut() != 1 {
			panic("unexpected argument")
		}
		_, err := h.Write([]byte(fnT.In(i).Out(0).String()))
		if err != nil {
			panic(err)
		}
	}
	return Args(h.Sum64())
}

func HashType[T any]() Type {
	h := fnv.New64()
	_, err := h.Write([]byte(reflect.TypeFor[T]().String()))
	if err != nil {
		panic(err)
	}
	return Type(h.Sum64())
}
