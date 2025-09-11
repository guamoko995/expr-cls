package env

import (
	"fmt"
	"math"
	"strings"

	"github.com/guamoko995/expr-cls/builder/makebuilder"
	"github.com/guamoko995/expr-cls/internal/hash"
)

func (env *Env) SetDefaultBinares() {
	defBinaryViaBuilderNoErr("or",
		makebuilder.FuncIn2(func(a, b bool) bool { return a || b }),
	)

	defBinaryViaBuilderNoErr("||",
		makebuilder.FuncIn2(func(a, b bool) bool { return a || b }),
	)

	defBinaryViaBuilderNoErr("and",
		makebuilder.FuncIn2(func(a, b bool) bool { return a && b }),
	)

	defBinaryViaBuilderNoErr("&&",
		makebuilder.FuncIn2(func(a, b bool) bool { return a && b }),
	)

	defBinaryViaBuilderNoErr("==",
		makebuilder.FuncIn2(func(a, b bool) bool { return a == b }),
		makebuilder.FuncIn2(func(a, b int) bool { return a == b }),
		makebuilder.FuncIn2(func(a, b string) bool { return a == b }),
	)

	defBinaryViaBuilderNoErr("!=",
		makebuilder.FuncIn2(func(a, b bool) bool { return a != b }),
		makebuilder.FuncIn2(func(a, b int) bool { return a != b }),
		makebuilder.FuncIn2(func(a, b string) bool { return a != b }),
	)

	defBinaryViaBuilderNoErr("<",
		makebuilder.FuncIn2(func(a, b float64) bool { return a < b }),
		makebuilder.FuncIn2(func(a int, b float64) bool { return float64(a) < b }),
		makebuilder.FuncIn2(func(a float64, b int) bool { return a < float64(b) }),
		makebuilder.FuncIn2(func(a, b int) bool { return a < b }),
		makebuilder.FuncIn2(func(a, b string) bool { return a < b }),
	)

	defBinaryViaBuilderNoErr(">",
		makebuilder.FuncIn2(func(a, b float64) bool { return a > b }),
		makebuilder.FuncIn2(func(a int, b float64) bool { return float64(a) > b }),
		makebuilder.FuncIn2(func(a float64, b int) bool { return a > float64(b) }),
		makebuilder.FuncIn2(func(a, b int) bool { return a > b }),
		makebuilder.FuncIn2(func(a, b string) bool { return a > b }),
	)

	defBinaryViaBuilderNoErr(">=",
		makebuilder.FuncIn2(func(a, b float64) bool { return a >= b }),
		makebuilder.FuncIn2(func(a int, b float64) bool { return float64(a) >= b }),
		makebuilder.FuncIn2(func(a float64, b int) bool { return a >= float64(b) }),
		makebuilder.FuncIn2(func(a, b int) bool { return a >= b }),
		makebuilder.FuncIn2(func(a, b string) bool { return a >= b }),
	)

	defBinaryViaBuilderNoErr("<=",
		makebuilder.FuncIn2(func(a, b float64) bool { return a <= b }),
		makebuilder.FuncIn2(func(a int, b float64) bool { return float64(a) <= b }),
		makebuilder.FuncIn2(func(a float64, b int) bool { return a <= float64(b) }),
		makebuilder.FuncIn2(func(a, b int) bool { return a <= b }),
		makebuilder.FuncIn2(func(a, b string) bool { return a <= b }),
	)

	defBinaryViaBuilderNoErr("+",
		makebuilder.FuncIn2(func(a, b int) int { return a + b }),
		makebuilder.FuncIn2(func(a float64, b float64) float64 { return a + b }),
		makebuilder.FuncIn2(func(a int, b float64) float64 { return float64(a) + b }),
		makebuilder.FuncIn2(func(a float64, b int) float64 { return a + float64(b) }),
		makebuilder.FuncIn2(func(a, b string) string { return a + b }),
	)

	defBinaryViaBuilderNoErr("-",
		makebuilder.FuncIn2(func(a, b int) int { return a - b }),
		makebuilder.FuncIn2(func(a float64, b float64) float64 { return a - b }),
		makebuilder.FuncIn2(func(a int, b float64) float64 { return float64(a) - b }),
		makebuilder.FuncIn2(func(a float64, b int) float64 { return a - float64(b) }),
	)

	defBinaryViaBuilderNoErr("*",
		makebuilder.FuncIn2(func(a, b int) int { return a * b }),
		makebuilder.FuncIn2(func(a float64, b float64) float64 { return a * b }),
		makebuilder.FuncIn2(func(a int, b float64) float64 { return float64(a) * b }),
		makebuilder.FuncIn2(func(a float64, b int) float64 { return a * float64(b) }),
		makebuilder.FuncIn2(func(a string, b int) string { return strings.Repeat(a, int(b)) }),
	)

	defBinaryViaBuilderNoErr("/",
		makebuilder.FuncIn2(func(a, b int) int { return a / b }),
		makebuilder.FuncIn2(func(a float64, b float64) float64 { return a / b }),
		makebuilder.FuncIn2(func(a int, b float64) float64 { return float64(a) / b }),
		makebuilder.FuncIn2(func(a float64, b int) float64 { return a / float64(b) }),
	)

	defBinaryViaBuilderNoErr("%",
		makebuilder.FuncIn2(func(a, b int) int { return a % b }),
	)

	defBinaryViaBuilderNoErr("**",
		makebuilder.FuncIn2(func(a, b int) int { return int(math.Pow(float64(a), float64(b))) }),
		makebuilder.FuncIn2(func(a float64, b float64) float64 { return math.Pow(a, b) }),
		makebuilder.FuncIn2(func(a int, b float64) float64 { return math.Pow(float64(a), b) }),
		makebuilder.FuncIn2(func(a float64, b int) float64 { return math.Pow(a, float64(b)) }),
	)

	defBinaryViaBuilderNoErr("^",
		makebuilder.FuncIn2(func(a, b int) int { return int(math.Pow(float64(a), float64(b))) }),
		makebuilder.FuncIn2(func(a float64, b float64) float64 { return math.Pow(a, b) }),
		makebuilder.FuncIn2(func(a int, b float64) float64 { return math.Pow(float64(a), b) }),
		makebuilder.FuncIn2(func(a float64, b int) float64 { return math.Pow(a, float64(b)) }),
	)
}

func defBinaryViaBuilderNoErr(token string, builder ...any) {
	if err := DefBinaryViaBuilder(token, Global, builder...); err != nil {
		panic(err)
	}
}

func DefBinaryViaBuilder(token string, env *Env, builder ...any) error {
	if _, exist := env.Binares[token]; !exist {
		return fmt.Errorf("operator %q is not supported", token)
	}

	for i := range builder {
		env.Binares[token][hash.HashArgsByBuilder(builder[i])] = builder[i]
	}
	return nil
}
