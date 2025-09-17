package defenv

import (
	"math"
	"strings"

	"github.com/guamoko995/expr-cls/env"
	"github.com/guamoko995/expr-cls/env/registrators"
)

// RegisterBinares registers binary operators in the environment.
func RegisterBinares(env *env.Enviroment) {
	env.RegisterBinaryNoErr("or",
		registrators.NewBinary(func(a, b bool) bool { return a || b }),
	)

	env.RegisterBinaryNoErr("||",
		registrators.NewBinary(func(a, b bool) bool { return a || b }),
	)

	env.RegisterBinaryNoErr("and",
		registrators.NewBinary(func(a, b bool) bool { return a && b }),
	)

	env.RegisterBinaryNoErr("&&",
		registrators.NewBinary(func(a, b bool) bool { return a && b }),
	)

	env.RegisterBinaryNoErr("==",
		registrators.NewBinary(func(a, b bool) bool { return a == b }),
		registrators.NewBinary(func(a, b int) bool { return a == b }),
		registrators.NewBinary(func(a, b string) bool { return a == b }),
	)

	env.RegisterBinaryNoErr("!=",
		registrators.NewBinary(func(a, b bool) bool { return a != b }),
		registrators.NewBinary(func(a, b int) bool { return a != b }),
		registrators.NewBinary(func(a, b string) bool { return a != b }),
	)

	env.RegisterBinaryNoErr("<",
		registrators.NewBinary(func(a, b float64) bool { return a < b }),
		registrators.NewBinary(func(a int, b float64) bool { return float64(a) < b }),
		registrators.NewBinary(func(a float64, b int) bool { return a < float64(b) }),
		registrators.NewBinary(func(a, b int) bool { return a < b }),
		registrators.NewBinary(func(a, b string) bool { return a < b }),
	)

	env.RegisterBinaryNoErr(">",
		registrators.NewBinary(func(a, b float64) bool { return a > b }),
		registrators.NewBinary(func(a int, b float64) bool { return float64(a) > b }),
		registrators.NewBinary(func(a float64, b int) bool { return a > float64(b) }),
		registrators.NewBinary(func(a, b int) bool { return a > b }),
		registrators.NewBinary(func(a, b string) bool { return a > b }),
	)

	env.RegisterBinaryNoErr(">=",
		registrators.NewBinary(func(a, b float64) bool { return a >= b }),
		registrators.NewBinary(func(a int, b float64) bool { return float64(a) >= b }),
		registrators.NewBinary(func(a float64, b int) bool { return a >= float64(b) }),
		registrators.NewBinary(func(a, b int) bool { return a >= b }),
		registrators.NewBinary(func(a, b string) bool { return a >= b }),
	)

	env.RegisterBinaryNoErr("<=",
		registrators.NewBinary(func(a, b float64) bool { return a <= b }),
		registrators.NewBinary(func(a int, b float64) bool { return float64(a) <= b }),
		registrators.NewBinary(func(a float64, b int) bool { return a <= float64(b) }),
		registrators.NewBinary(func(a, b int) bool { return a <= b }),
		registrators.NewBinary(func(a, b string) bool { return a <= b }),
	)

	env.RegisterBinaryNoErr("+",
		registrators.NewBinary(func(a, b int) int { return a + b }),
		registrators.NewBinary(func(a float64, b float64) float64 { return a + b }),
		registrators.NewBinary(func(a int, b float64) float64 { return float64(a) + b }),
		registrators.NewBinary(func(a float64, b int) float64 { return a + float64(b) }),
		registrators.NewBinary(func(a, b string) string { return a + b }),
	)

	env.RegisterBinaryNoErr("-",
		registrators.NewBinary(func(a, b int) int { return a - b }),
		registrators.NewBinary(func(a float64, b float64) float64 { return a - b }),
		registrators.NewBinary(func(a int, b float64) float64 { return float64(a) - b }),
		registrators.NewBinary(func(a float64, b int) float64 { return a - float64(b) }),
	)

	env.RegisterBinaryNoErr("*",
		registrators.NewBinary(func(a, b int) int { return a * b }),
		registrators.NewBinary(func(a float64, b float64) float64 { return a * b }),
		registrators.NewBinary(func(a int, b float64) float64 { return float64(a) * b }),
		registrators.NewBinary(func(a float64, b int) float64 { return a * float64(b) }),
		registrators.NewBinary(func(a string, b int) string { return strings.Repeat(a, int(b)) }),
	)

	env.RegisterBinaryNoErr("/",
		registrators.NewBinary(func(a, b int) int { return a / b }),
		registrators.NewBinary(func(a float64, b float64) float64 { return a / b }),
		registrators.NewBinary(func(a int, b float64) float64 { return float64(a) / b }),
		registrators.NewBinary(func(a float64, b int) float64 { return a / float64(b) }),
	)

	env.RegisterBinaryNoErr("%",
		registrators.NewBinary(func(a, b int) int { return a % b }),
	)

	env.RegisterBinaryNoErr("**",
		registrators.NewBinary(func(a, b int) int { return int(math.Pow(float64(a), float64(b))) }),
		registrators.NewBinary(func(a float64, b float64) float64 { return math.Pow(a, b) }),
		registrators.NewBinary(func(a int, b float64) float64 { return math.Pow(float64(a), b) }),
		registrators.NewBinary(func(a float64, b int) float64 { return math.Pow(a, float64(b)) }),
	)

	env.RegisterBinaryNoErr("^",
		registrators.NewBinary(func(a, b int) int { return int(math.Pow(float64(a), float64(b))) }),
		registrators.NewBinary(func(a float64, b float64) float64 { return math.Pow(a, b) }),
		registrators.NewBinary(func(a int, b float64) float64 { return math.Pow(float64(a), b) }),
		registrators.NewBinary(func(a float64, b int) float64 { return math.Pow(a, float64(b)) }),
	)
}
