package hashsum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	var (
		f1 func(byte, int64) float64
		b1 func(func() byte, func() int64) func() float64

		f2 func(byte, int64) string
		b2 func(func() byte, func() int64) func() string

		f3 func(int64, byte) string
		b3 func(func() int64, func() byte) func() string

		f4 func(byte) string
		b4 func(func() byte) func() string
	)

	t.Run("equivalence of hashes by function and by builder", func(t *testing.T) {
		require.Equal(t, HashArgsByFunc(f1), HashArgsByBuilder(b1))
		require.Equal(t, HashArgsByFunc(f2), HashArgsByBuilder(b2))
		require.Equal(t, HashArgsByFunc(f3), HashArgsByBuilder(b3))
		require.Equal(t, HashArgsByFunc(f4), HashArgsByBuilder(b4))
	})

	t.Run("equivalence of hashes for functions with the same number and types of arguments", func(t *testing.T) {
		require.Equal(t, HashArgsByFunc(f1), HashArgsByFunc(f2))
		require.Equal(t, HashArgsByBuilder(b1), HashArgsByBuilder(b2))
	})

	t.Run("different hashes for functions with different numbers or types of arguments", func(t *testing.T) {
		require.NotEqual(t, HashArgsByFunc(f1), HashArgsByFunc(f3))
		require.NotEqual(t, HashArgsByBuilder(b1), HashArgsByBuilder(b3))

		require.NotEqual(t, HashArgsByFunc(f1), HashArgsByFunc(f4))
		require.NotEqual(t, HashArgsByBuilder(b1), HashArgsByBuilder(b4))

		require.NotEqual(t, HashArgsByFunc(f3), HashArgsByFunc(f4))
		require.NotEqual(t, HashArgsByBuilder(b3), HashArgsByBuilder(b4))
	})
}
