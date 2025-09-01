package checker_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/guamoko995/expr-cls/internal/testify/require"

	"github.com/guamoko995/expr-cls/checker"
	"github.com/guamoko995/expr-cls/test/mock"
)

func TestTypedFuncIndex(t *testing.T) {
	fn := func() time.Duration {
		return 1 * time.Second
	}
	index, ok := checker.TypedFuncIndex(reflect.TypeOf(fn), false)
	require.True(t, ok)
	require.Equal(t, 1, index)
}

func TestTypedFuncIndex_excludes_named_functions(t *testing.T) {
	var fn mock.MyFunc

	_, ok := checker.TypedFuncIndex(reflect.TypeOf(fn), false)
	require.False(t, ok)
}
