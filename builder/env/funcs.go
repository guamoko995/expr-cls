package env

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/guamoko995/expr-cls/builder/base"
	"github.com/guamoko995/expr-cls/internal/hash"
)

func (env *Env) SetDefaultFuncs() {
	defFuncViaBuilderNoErr("trim",
		base.MakeFuncIn1Builder(strings.TrimSpace),
		base.MakeFuncIn2Builder(strings.Trim),
	)

	defFuncViaBuilderNoErr("trimPrefix",
		base.MakeFuncIn1Builder(func(s string) string { return strings.TrimPrefix(s, " ") }),
		base.MakeFuncIn2Builder(strings.TrimPrefix),
	)

	defFuncViaBuilderNoErr("trimSuffix",
		base.MakeFuncIn1Builder(func(s string) string { return strings.TrimSuffix(s, " ") }),
		base.MakeFuncIn2Builder(strings.TrimSuffix),
	)

	defFuncViaBuilderNoErr("lower",
		base.MakeFuncIn1Builder(strings.ToLower),
	)

	defFuncViaBuilderNoErr("split",
		base.MakeFuncIn2Builder(strings.Split),
		base.MakeFuncIn3Builder(strings.SplitN),
	)

	defFuncViaBuilderNoErr("splitAfter",
		base.MakeFuncIn2Builder(strings.SplitAfter),
		base.MakeFuncIn3Builder(strings.SplitAfterN),
	)

	defFuncViaBuilderNoErr("replace",
		base.MakeFuncIn3Builder(strings.ReplaceAll),
		base.MakeFuncIn4Builder(strings.Replace),
	)

	defFuncViaBuilderNoErr("repeat",
		base.MakeFuncIn2Builder(strings.Repeat),
	)

	defFuncViaBuilderNoErr("join",
		base.MakeFuncIn1Builder(func(elems []string) string { return strings.Join(elems, "") }),
		base.MakeFuncIn2Builder(strings.Join),
	)

	defFuncViaBuilderNoErr("lastIndexOf",
		base.MakeFuncIn2Builder(strings.LastIndex),
	)

	defFuncViaBuilderNoErr("hasPrefix",
		base.MakeFuncIn2Builder(strings.HasPrefix),
	)

	defFuncViaBuilderNoErr("hasSuffix",
		base.MakeFuncIn2Builder(strings.HasSuffix),
	)

	defFuncViaBuilderNoErr("max",
		base.MakeFuncIn2Builder(math.Max),
		base.MakeFuncIn2Builder(func(x int, y float64) float64 { return math.Max(float64(x), y) }),
		base.MakeFuncIn2Builder(func(x float64, y int) float64 { return math.Max(x, float64(y)) }),
		base.MakeFuncIn2Builder(func(x, y int) int { return int(math.Max(float64(x), float64(y))) }),
	)

	defFuncViaBuilderNoErr("min",
		base.MakeFuncIn2Builder(math.Min),
		base.MakeFuncIn2Builder(func(x int, y float64) float64 { return math.Min(float64(x), y) }),
		base.MakeFuncIn2Builder(func(x float64, y int) float64 { return math.Min(x, float64(y)) }),
		base.MakeFuncIn2Builder(func(x, y int) int { return int(math.Min(float64(x), float64(y))) }),
	)
}

func defFuncViaBuilderNoErr(token string, builder ...base.Builder) {
	for i := range builder {
		if err := DefFuncViaBuilder(token, Global, builder...); err != nil {
			panic(fmt.Errorf("failed def %d (starting from 0) builder: %w", i, err))
		}
	}
}

func DefFuncViaBuilder(token string, env *Env, builder ...base.Builder) error {
	if _, exist := env.UnaryBuilders[token]; exist {
		return errors.New("token is reserved for operators")
	}

	if _, exist := env.BinaryBuilders[token]; exist {
		return errors.New("token is reserved for operators")
	}

	if _, exist := env.Consts[token]; exist {
		return errors.New("token is reserved by const")
	}

	if _, exist := env.FunctionBuilders[token]; !exist {
		env.FunctionBuilders[token] = make(map[hash.Args]base.Builder)
	}

	for i := range builder {
		env.FunctionBuilders[token][hash.HashArgsByBuilder(builder[i])] = builder[i]
	}

	return nil
}
