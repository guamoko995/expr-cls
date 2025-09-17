package defenv

import (
	"math"
	"strings"

	"github.com/guamoko995/expr-cls/env"
	"github.com/guamoko995/expr-cls/env/registrators"
)

// RegisterFuncs registers default funcs in the current environment.
func RegisterFuncs(env *env.Enviroment) {
	env.RegisterFuncNoErr("trim",
		registrators.NewFuncIn1(strings.TrimSpace),
		registrators.NewFuncIn2(strings.Trim),
	)

	env.RegisterFuncNoErr("trimPrefix",
		registrators.NewFuncIn1(func(s string) string { return strings.TrimPrefix(s, " ") }),
		registrators.NewFuncIn2(strings.TrimPrefix),
	)

	env.RegisterFuncNoErr("trimSuffix",
		registrators.NewFuncIn1(func(s string) string { return strings.TrimSuffix(s, " ") }),
		registrators.NewFuncIn2(strings.TrimSuffix),
	)

	env.RegisterFuncNoErr("lower",
		registrators.NewFuncIn1(strings.ToLower),
	)

	env.RegisterFuncNoErr("split",
		registrators.NewFuncIn2(strings.Split),
		registrators.NewFuncIn3(strings.SplitN),
	)

	env.RegisterFuncNoErr("splitAfter",
		registrators.NewFuncIn2(strings.SplitAfter),
		registrators.NewFuncIn3(strings.SplitAfterN),
	)

	env.RegisterFuncNoErr("replace",
		registrators.NewFuncIn3(strings.ReplaceAll),
		registrators.NewFuncIn4(strings.Replace),
	)

	env.RegisterFuncNoErr("repeat",
		registrators.NewFuncIn2(strings.Repeat),
	)

	env.RegisterFuncNoErr("join",
		registrators.NewFuncIn1(func(elems []string) string { return strings.Join(elems, "") }),
		registrators.NewFuncIn2(strings.Join),
	)

	env.RegisterFuncNoErr("lastIndexOf",
		registrators.NewFuncIn2(strings.LastIndex),
	)

	env.RegisterFuncNoErr("hasPrefix",
		registrators.NewFuncIn2(strings.HasPrefix),
	)

	env.RegisterFuncNoErr("hasSuffix",
		registrators.NewFuncIn2(strings.HasSuffix),
	)

	env.RegisterFuncNoErr("max",
		registrators.NewFuncIn2(math.Max),
		registrators.NewFuncIn2(func(x int, y float64) float64 { return math.Max(float64(x), y) }),
		registrators.NewFuncIn2(func(x float64, y int) float64 { return math.Max(x, float64(y)) }),
		registrators.NewFuncIn2(func(x, y int) int { return int(math.Max(float64(x), float64(y))) }),
	)

	env.RegisterFuncNoErr("min",
		registrators.NewFuncIn2(math.Min),
		registrators.NewFuncIn2(func(x int, y float64) float64 { return math.Min(float64(x), y) }),
		registrators.NewFuncIn2(func(x float64, y int) float64 { return math.Min(x, float64(y)) }),
		registrators.NewFuncIn2(func(x, y int) int { return int(math.Min(float64(x), float64(y))) }),
	)
}
