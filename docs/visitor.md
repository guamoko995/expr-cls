# Visitor

Expr provides an interface to traverse the <span title="Abstract Syntax Tree" style={{borderBottom: "1px dotted currentColor"}}>AST</span> of the expression before the compilation.
The `Visitor` interface allows you to collect information about the expression, modify the expression, or even generate
a new expression.

Let's start with an [ast.Visitor](https://pkg.go.dev/github.com/guamoko995/expr-cls/ast#Visitor) implementation which will 
collect all variables used in the expression.

Visitor must implement a single method `Visit(*ast.Node)`, which will be called for each node in the AST.

```go
type Visitor struct {
    Identifiers []string
}

func (v *Visitor) Visit(node *ast.Node) {
    if n, ok := (*node).(*ast.IdentifierNode); ok {
        v.Identifiers = append(v.Identifiers, n.Value)
    }
}
```

Full list of available AST nodes can be found in the [ast](https://pkg.go.dev/github.com/guamoko995/expr-cls/ast) documentation.

Let's parse the expression and use [ast.Walk](https://pkg.go.dev/github.com/guamoko995/expr-cls/ast#Walk) to traverse the AST:

```go
tree, err := parser.Parse(`foo + bar`)
if err != nil {
    panic(err)
}

v := &Visitor{}
// highlight-next-line
ast.Walk(&tree.Node, v)

fmt.Println(v.Identifiers) // [foo, bar]
```

:::note

Although it is possible to access the AST of compiled program, it may be already be modified by patchers, optimizers, etc.

```go
program, err := expr.Compile(`foo + bar`)
if err != nil {
    panic(err)
}

// highlight-next-line
node := program.Node()

v := &Visitor{}
ast.Walk(&node, v)
```

:::
