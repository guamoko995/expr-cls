# expr-cls

**expr-cls** is a minimal MVP implementation of an expression compilation and evaluation engine, supporting custom variables, operators, and lazy evaluation.  
**Key advantage:** evaluation speed is about 17x faster than [expr-lang/expr](https://github.com/expr-lang/expr) (see Benchmark section below).

This fork introduces a new architecture, making it easy to extend functions and variable sources, and overload operators.

## Features (MVP Stage)

- **Custom variable sources:** Use struct fields as variables in expressions.
- **Operator overloading:** Arithmetic and logical operators can be overloaded for custom types via builder interface (parser limitation: operators cannot be extended, only overloaded).
- **Functions:** Built-in string and math functions, extensible with your own.
- **Constants:** Built-in mathematical constants (`pi`, `phi`).
- **Lazy evaluation:** Expressions compile into Go functions.
- **Extensible environment:** Register new variable types, functions, and constants.

## Quick Example

```go
package main

import (
    "fmt"
    "github.com/guamoko995/expr-cls/expr"
)

func main() {
    // Define input structure
    type InputData struct {
        A int
        B int
    }

    // Register the struct type as a variable source
    expr.RegisterSource[InputData]()

    // Compile the expression
    fn, err := expr.Compile[InputData, int]("3 + A * -B")
    if err != nil {
        panic(err)
    }

    // Provide input data
    input := InputData{A: 7, B: 10}

    // Evaluate the expression
    result := fn(input)

    fmt.Println(result) // Output: -67
}
```

## Benchmark

Below is a comparison between `expr-cls` and [expr-lang/expr](https://github.com/expr-lang/expr):

```
Benchmark/concept-12         134553903        8.947 ns/op        0 B/op   0 allocs/op
Benchmark/expr-12            7743223        156.6 ns/op        136 B/op 6 allocs/op
PASS
ok      github.com/guamoko995/expr-cls/proof_of_concept   2.422s
```

**Explanation:**  
`expr-cls` compiles expressions directly to lazy Go functions using struct variable sources, which is extremely fast and zero-alloc for supported types. The expr-lang/expr library is, for now, more mature and universal, but slower and requires memory allocation.

## Architecture Overview

- **AST → Builder → LazyFunc:** Expressions are parsed to AST, then compiled to lazy-evaluated functions via builder interfaces.
- **Environment (`env`):** Holds operator/function/constant/variable builders. Easily extendable.
- **Operator overloading:** Operators can be overloaded for custom types, but the set of operators is fixed by the parser.
- **Functions:** Registered via builder interfaces (`base.Builder`). Supports unary, binary, and multi-argument functions.
- **Variables:** Use struct fields (after registration) as variable sources in expressions.

### Extending the Environment

You can register your own variable types, functions, and constants:

```go
import "github.com/guamoko995/expr-cls/builder/env"

// Register new variable type
env.RegVarType[MyType](env.Global)

// Overload operator for custom types
env.DefUnaryViaBuilder("myop", env.Global, myUnaryBuilder)

// Register new function
env.DefFuncViaBuilder("myfunc", env.Global, myFuncBuilder)

// Register new constant
env.RegConst("e", env.Global, math.E)
```

## Built-in Operators & Functions

- Arithmetic: `+`, `-`, `*`, `/`, `%`, `**`, `^`
- Logical: `and`, `or`, `not`, `!`
- Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
- String: `trim`, `trimPrefix`, `trimSuffix`, `lower`, `split`, `splitAfter`, `replace`, `repeat`, `join`, `lastIndexOf`, `hasPrefix`, `hasSuffix`
- Math: `max`, `min`
- Constants: `pi`, `phi`

## MVP Limitations

- Only struct variable sources are supported (for now).
- Operators cannot be extended; only overloaded for supported types.
- Some original expr-lang/expr features and builtins are disabled or commented (see code for details).
- Focus on performance and simplicity for the MVP.

## Reference

- [Proof of concept & benchmarks](proof_of_concept/poc_bench_test.go)
- [Usage example](expr_test.go)

---

This README reflects the new MVP architecture, usage patterns, and extensibility of **expr-cls**. Feel free to contribute or extend!
