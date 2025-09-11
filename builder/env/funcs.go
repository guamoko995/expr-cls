package env

import (
	"fmt"
	"math"
	"strings"

	"github.com/guamoko995/expr-cls/builder/makebuilder"
	"github.com/guamoko995/expr-cls/internal/hash"
)

func (env *Env) SetDefaultFuncs() {
	defFuncViaBuilderNoErr("trim",
		makebuilder.FuncIn1(strings.TrimSpace),
		makebuilder.FuncIn2(strings.Trim),
	)

	defFuncViaBuilderNoErr("trimPrefix",
		makebuilder.FuncIn1(func(s string) string { return strings.TrimPrefix(s, " ") }),
		makebuilder.FuncIn2(strings.TrimPrefix),
	)

	defFuncViaBuilderNoErr("trimSuffix",
		makebuilder.FuncIn1(func(s string) string { return strings.TrimSuffix(s, " ") }),
		makebuilder.FuncIn2(strings.TrimSuffix),
	)

	defFuncViaBuilderNoErr("lower",
		makebuilder.FuncIn1(strings.ToLower),
	)

	defFuncViaBuilderNoErr("split",
		makebuilder.FuncIn2(strings.Split),
		makebuilder.FuncIn3(strings.SplitN),
	)

	defFuncViaBuilderNoErr("splitAfter",
		makebuilder.FuncIn2(strings.SplitAfter),
		makebuilder.FuncIn3(strings.SplitAfterN),
	)

	defFuncViaBuilderNoErr("replace",
		makebuilder.FuncIn3(strings.ReplaceAll),
		makebuilder.FuncIn4(strings.Replace),
	)

	defFuncViaBuilderNoErr("repeat",
		makebuilder.FuncIn2(strings.Repeat),
	)

	defFuncViaBuilderNoErr("join",
		makebuilder.FuncIn1(func(elems []string) string { return strings.Join(elems, "") }),
		makebuilder.FuncIn2(strings.Join),
	)

	defFuncViaBuilderNoErr("lastIndexOf",
		makebuilder.FuncIn2(strings.LastIndex),
	)

	defFuncViaBuilderNoErr("hasPrefix",
		makebuilder.FuncIn2(strings.HasPrefix),
	)

	defFuncViaBuilderNoErr("hasSuffix",
		makebuilder.FuncIn2(strings.HasSuffix),
	)

	defFuncViaBuilderNoErr("max",
		makebuilder.FuncIn2(math.Max),
		makebuilder.FuncIn2(func(x int, y float64) float64 { return math.Max(float64(x), y) }),
		makebuilder.FuncIn2(func(x float64, y int) float64 { return math.Max(x, float64(y)) }),
		makebuilder.FuncIn2(func(x, y int) int { return int(math.Max(float64(x), float64(y))) }),
	)

	defFuncViaBuilderNoErr("min",
		makebuilder.FuncIn2(math.Min),
		makebuilder.FuncIn2(func(x int, y float64) float64 { return math.Min(float64(x), y) }),
		makebuilder.FuncIn2(func(x float64, y int) float64 { return math.Min(x, float64(y)) }),
		makebuilder.FuncIn2(func(x, y int) int { return int(math.Min(float64(x), float64(y))) }),
	)
}

func defFuncViaBuilderNoErr(token string, builder ...any) {
	for i := range builder {
		if err := DefFuncViaBuilder(token, Global, builder); err != nil {
			panic(fmt.Errorf("failed def %d (starting from 0) builder: %w", i, err))
		}
	}
}

func DefFuncViaBuilder(token string, env *Env, builder ...any) error {
	if _, exist := env.Unares[token]; exist {
		return fmt.Errorf("token %q is reserved for operators", token)
	}

	if _, exist := env.Binares[token]; exist {
		return fmt.Errorf("token %q is reserved for operators", token)
	}

	if _, exist := env.Functions[token]; !exist {
		env.Functions[token] = make(map[hash.Args]any)
	}

	for i := range builder {
		env.Functions[token][hash.HashArgsByBuilder(builder[i])] = builder[i]
	}

	return nil
}
