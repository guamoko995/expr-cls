# expr-cls

**expr-cls** is a minimal, experimental implementation of a high-performance string expression compiler and runtime for Go.  
The core idea is to build expressions as chains of strictly-typed Go closures (functions), rather than interpreting bytecode on a virtual machine.  
This architecture enables extremely fast compilation and execution, zero allocations during evaluation, and a flexible, extensible environment system.

---

## Key Concepts

**Environment:**  
At compile time, you define an environment describing:
- Supported unary and binary operators for the parser (not implemented yet; currently a fixed set is used).
- Complex constructs for the parser (e.g., ternary operator, arbitrary typed literals; planned).
- Overload registrations for unary and binary operators.
- Overload registrations for functions.
- Constant registrations.
- Variable type registrations.
- Variable source type registrations (struct fields as variables).

**Extensibility:**  
You can create separate packages with ready-made environments for domain-specific tasks (e.g., matrices, complex numbers, statistics, geospatial, etc).
See [`example`](https://github.com/guamoko995/expr-cls/tree/master/tests/example) for a current example of environment definition and usage.

**Performance:**  
Expressions compile and execute extremely quickly at runtime.  
Compiled expressions are strictly typed.

---

## Usage

```go
package example_test

import (
	"fmt"

	exprcls "github.com/guamoko995/expr-cls"

	// Using the example environment
	_ "github.com/guamoko995/expr-cls/tests/example/def_env"
)

// CompileAndCalcExample demonstrates how to compile and evaluate expressions
// using the expr-cls package.
func Example() {
	// Define a data structure containing input variables for our expression.
	type InputData struct {
		A int
		B int
	}

	// Parse and compile the expression "3 + A * B".
	prog, err := exprcls.Compile[InputData, float64]("3 + A * -B + sin(pi/2)")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Provide input data for evaluating the expression.
	input := InputData{A: 7, B: 10}

	// Evaluate the expression using the provided input data.
	result := prog(input)

	// Print the computed result.
	fmt.Println(result)

	// Output: -66
}
```

---

## Features

- Struct fields as variables in expressions (via registration).
- Operator overloading for custom types (operator set is fixed by parser, not extendable yet).
- Function overloading and custom function registration.
- Constant registration.
- Strict typing of compiled expressions.
- Extremely fast compilation and evaluation (see benchmarks below).
- MVP: Some features and built-ins from original [expr-lang/expr](https://github.com/expr-lang/expr) are disabled or not implemented.

---

## Benchmarks

You can find the benchmark source in [`tests/benchmarks`](https://github.com/guamoko995/expr-cls/tree/master/tests/benchmarks).

### Compilation Speed

Measures the time and resource usage to compile an expression ([benchmark code](https://github.com/guamoko995/expr-cls/tree/master/tests/benchmarks)):

```
goos: linux
goarch: amd64
pkg: github.com/guamoko995/expr-cls/tests/benchmarks
cpu: AMD Ryzen 5 5600H with Radeon Graphics         
BenchmarkCompile/expr-cls-12         	  644955	      1730 ns/op	    1344 B/op	      31 allocs/op
BenchmarkCompile/expr-12             	  110992	     10489 ns/op	   10342 B/op	      75 allocs/op
PASS
ok  	github.com/guamoko995/expr-cls/tests/benchmarks	2.285s
```
**expr-cls compiles more than 6x faster and with 8x less memory allocation than expr-lang/expr.**

---

### Evaluation Speed

Measures the time and resource usage to repeatedly evaluate a compiled expression ([benchmark code](https://github.com/guamoko995/expr-cls/tree/master/tests/benchmarks)):

```
expression: "X+(6*Y)"
params:
	X=3
	Y=5
expr-cls result: 33
expr result: 33

goos: linux
goarch: amd64
pkg: github.com/guamoko995/expr-cls/tests/benchmarks
cpu: AMD Ryzen 5 5600H with Radeon Graphics         
BenchmarkСalc/expr-cls-12          	134900941	         8.885 ns/op	       0 B/op	       0 allocs/op
BenchmarkСalc/expr-12              	 7426069	       162.8 ns/op	     136 B/op	       6 allocs/op
PASS
ok  	github.com/guamoko995/expr-cls/tests/benchmarks	2.413s
```
**expr-cls evaluates expressions ~18x faster and with zero allocations.**

---

## Architecture

- **Parsing:** Expressions are parsed into AST nodes. The set of operators/constructs is currently fixed.
- **Environment:** Holds operator/function/constant/variable builders. Easily extended. See [`example`](https://github.com/guamoko995/expr-cls/tree/master/tests/example/) for practical setup.
- **Building:** AST is compiled into a closure chain (Go functions), not bytecode.
- **Evaluation:** The compiled closure chain receives strictly-typed input (struct) and returns strictly-typed output.

---

## Extending expr-cls

To see how to register types, constants, functions, and operator overloads, refer to the [`example`](https://github.com/guamoko995/expr-cls/tree/master/tests/example/).

---

## Limitations and Roadmap

- Only struct variable sources supported (for now).
- Operators cannot be extended (only overloaded).
- Some complex constructs (e.g. ternary operator, custom literals) are planned but not implemented.
- Error handling and reporting are minimal.
- More tests and environments are planned.
- Parser extensibility: declarative operator/construct definition is a future goal.

---

**This README summarizes the ideas and experimental architecture of expr-cls.  
Feedback and contributions are welcome.**