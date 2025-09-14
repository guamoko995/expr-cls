package env

import (
	"fmt"
	"math"
	"strings"

	"github.com/guamoko995/expr-cls/builder/base"
	"github.com/guamoko995/expr-cls/internal/hash"
)

func (env *Env) SetDefaultBinares() {
	defBinaryViaBuilderNoErr("or",
		base.MakeFuncIn2Builder(func(a, b bool) bool { return a || b }),
	)

	defBinaryViaBuilderNoErr("||",
		base.MakeFuncIn2Builder(func(a, b bool) bool { return a || b }),
	)

	defBinaryViaBuilderNoErr("and",
		base.MakeFuncIn2Builder(func(a, b bool) bool { return a && b }),
	)

	defBinaryViaBuilderNoErr("&&",
		base.MakeFuncIn2Builder(func(a, b bool) bool { return a && b }),
	)

	defBinaryViaBuilderNoErr("==",
		base.MakeFuncIn2Builder(func(a, b bool) bool { return a == b }),
		base.MakeFuncIn2Builder(func(a, b int) bool { return a == b }),
		base.MakeFuncIn2Builder(func(a, b string) bool { return a == b }),
	)

	defBinaryViaBuilderNoErr("!=",
		base.MakeFuncIn2Builder(func(a, b bool) bool { return a != b }),
		base.MakeFuncIn2Builder(func(a, b int) bool { return a != b }),
		base.MakeFuncIn2Builder(func(a, b string) bool { return a != b }),
	)

	defBinaryViaBuilderNoErr("<",
		base.MakeFuncIn2Builder(func(a, b float64) bool { return a < b }),
		base.MakeFuncIn2Builder(func(a int, b float64) bool { return float64(a) < b }),
		base.MakeFuncIn2Builder(func(a float64, b int) bool { return a < float64(b) }),
		base.MakeFuncIn2Builder(func(a, b int) bool { return a < b }),
		base.MakeFuncIn2Builder(func(a, b string) bool { return a < b }),
	)

	defBinaryViaBuilderNoErr(">",
		base.MakeFuncIn2Builder(func(a, b float64) bool { return a > b }),
		base.MakeFuncIn2Builder(func(a int, b float64) bool { return float64(a) > b }),
		base.MakeFuncIn2Builder(func(a float64, b int) bool { return a > float64(b) }),
		base.MakeFuncIn2Builder(func(a, b int) bool { return a > b }),
		base.MakeFuncIn2Builder(func(a, b string) bool { return a > b }),
	)

	defBinaryViaBuilderNoErr(">=",
		base.MakeFuncIn2Builder(func(a, b float64) bool { return a >= b }),
		base.MakeFuncIn2Builder(func(a int, b float64) bool { return float64(a) >= b }),
		base.MakeFuncIn2Builder(func(a float64, b int) bool { return a >= float64(b) }),
		base.MakeFuncIn2Builder(func(a, b int) bool { return a >= b }),
		base.MakeFuncIn2Builder(func(a, b string) bool { return a >= b }),
	)

	defBinaryViaBuilderNoErr("<=",
		base.MakeFuncIn2Builder(func(a, b float64) bool { return a <= b }),
		base.MakeFuncIn2Builder(func(a int, b float64) bool { return float64(a) <= b }),
		base.MakeFuncIn2Builder(func(a float64, b int) bool { return a <= float64(b) }),
		base.MakeFuncIn2Builder(func(a, b int) bool { return a <= b }),
		base.MakeFuncIn2Builder(func(a, b string) bool { return a <= b }),
	)

	defBinaryViaBuilderNoErr("+",
		base.MakeFuncIn2Builder(func(a, b int) int { return a + b }),
		base.MakeFuncIn2Builder(func(a float64, b float64) float64 { return a + b }),
		base.MakeFuncIn2Builder(func(a int, b float64) float64 { return float64(a) + b }),
		base.MakeFuncIn2Builder(func(a float64, b int) float64 { return a + float64(b) }),
		base.MakeFuncIn2Builder(func(a, b string) string { return a + b }),
	)

	defBinaryViaBuilderNoErr("-",
		base.MakeFuncIn2Builder(func(a, b int) int { return a - b }),
		base.MakeFuncIn2Builder(func(a float64, b float64) float64 { return a - b }),
		base.MakeFuncIn2Builder(func(a int, b float64) float64 { return float64(a) - b }),
		base.MakeFuncIn2Builder(func(a float64, b int) float64 { return a - float64(b) }),
	)

	defBinaryViaBuilderNoErr("*",
		base.MakeFuncIn2Builder(func(a, b int) int { return a * b }),
		base.MakeFuncIn2Builder(func(a float64, b float64) float64 { return a * b }),
		base.MakeFuncIn2Builder(func(a int, b float64) float64 { return float64(a) * b }),
		base.MakeFuncIn2Builder(func(a float64, b int) float64 { return a * float64(b) }),
		base.MakeFuncIn2Builder(func(a string, b int) string { return strings.Repeat(a, int(b)) }),
	)

	defBinaryViaBuilderNoErr("/",
		base.MakeFuncIn2Builder(func(a, b int) int { return a / b }),
		base.MakeFuncIn2Builder(func(a float64, b float64) float64 { return a / b }),
		base.MakeFuncIn2Builder(func(a int, b float64) float64 { return float64(a) / b }),
		base.MakeFuncIn2Builder(func(a float64, b int) float64 { return a / float64(b) }),
	)

	defBinaryViaBuilderNoErr("%",
		base.MakeFuncIn2Builder(func(a, b int) int { return a % b }),
	)

	defBinaryViaBuilderNoErr("**",
		base.MakeFuncIn2Builder(func(a, b int) int { return int(math.Pow(float64(a), float64(b))) }),
		base.MakeFuncIn2Builder(func(a float64, b float64) float64 { return math.Pow(a, b) }),
		base.MakeFuncIn2Builder(func(a int, b float64) float64 { return math.Pow(float64(a), b) }),
		base.MakeFuncIn2Builder(func(a float64, b int) float64 { return math.Pow(a, float64(b)) }),
	)

	defBinaryViaBuilderNoErr("^",
		base.MakeFuncIn2Builder(func(a, b int) int { return int(math.Pow(float64(a), float64(b))) }),
		base.MakeFuncIn2Builder(func(a float64, b float64) float64 { return math.Pow(a, b) }),
		base.MakeFuncIn2Builder(func(a int, b float64) float64 { return math.Pow(float64(a), b) }),
		base.MakeFuncIn2Builder(func(a float64, b int) float64 { return math.Pow(a, float64(b)) }),
	)
}

func defBinaryViaBuilderNoErr(token string, builder ...base.Builder) {
	if err := DefBinaryViaBuilder(token, Global, builder...); err != nil {
		panic(err)
	}
}

func DefBinaryViaBuilder(token string, env *Env, builder ...base.Builder) error {
	if _, exist := env.BinaryBuilders[token]; !exist {
		return fmt.Errorf("operator %q is not supported", token)
	}

	for i := range builder {
		env.BinaryBuilders[token][hash.HashArgsByBuilder(builder[i])] = builder[i]
	}
	return nil
}
