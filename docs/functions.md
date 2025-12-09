
## Functions

Functions (and only functions) execute code.

### Function Statements

Functions can be created at the top of current scope with statements using the `func` keyword.

### Function Literals

Functions can be created at the current scope with literals using the `->` symbols.

### Environment

A function environment starts with variables from input and output parameters. Output variables are assigned zero values unless they are assigned values at the top of function scope.

#### Additional function-scope variables

At each scope within the function body, new variables must be declared at the top of scope.

```
func average(nums array<float>) (avg float) {
    var total float

    nums.each->(x float){
        total = total + x
    }

    avg = total / float(nums.len())
}
```

Because the environment is derived from parameter names, there cannot be overlapping input and output parameter names.

### Signature

Function signature is the sequence of input and output parameter types. If the sequence of either input types or output types are different between two functions, then they will have different signatures.

### Parameters

A function parameter has name and type. 

### Overloading

Functions can have the same name as long as the signature is different.

### Termination

Functions terminate either at the end of the function body or when a return statement is executed. Output parameter values must be assigned before termination or they will be left at their zero values.

### Parameter-Argument Binding

Input and output parameters can be bound to arguments before the function is called. Binding produces a new function without the parameters that were bound.

Binding an input parameter:
```
func repeat(s string, n nnint) {
    n.times->{
        std.out << s
    }
}

// bind input with ordinal arg
var repeat5 repeat$(_ 5)

// ok ok ok ok ok
repeat5("ok")

// no no no no no
repeat5("no")
```

Binding an output parameter:
```
func lookupName(
    book AddressBook
    name string)(
    success bool
    address string) {
    // TODO
}

var found bool
var addr string

// bind outputs with ordinal args
var lookup lookupName$()(found addr)

["Uriah Caw" "Caleb Plummer"].each->(name string) {
    if book.lookup(); found {
        std.out << "he lives at #{addr}"
    }
}
```

### Function Call Arguments

Call arguments can be provided in multiple ways:
* Ordinal-only
* Keyword-only
* Mixed Ordinal/Keyword

Some examples that demonstrate the calling conventions, using the function below.

```
func line(x, slope, intercept float) (y float) {
    y = x * slope + intercept
}
```

#### Ordinal Arguments

```
func main() {
    const m 2.5
    const c -0.7
    var ys array<float>

    for x in [-1.0...1.0] {
        var y line(x, m, c)
        
        ys.push(y)
    }
}
```

#### Auto-Keyword Arguments

```
func main() {
    const slope 2.5
    const intercept -0.7
    var ys array<float>

    for x in [-1.0...1.0] {
        var y line((x, m, c))
        
        ys.push(y)
    }
}
```

#### Mixed-Keyword Arguments

```
func main() {
    const m 2.5
    const c -0.7
    var ys array<float>

    for x in [-1.0...1.0] {
        var y line((x, slope: m, intercept: c))
        
        ys.push(y)
    }
}
```

#### Method Call Syntax

```
func main() {
    const m 2.5
    const c -0.7
    var ys array<float>

    for x in [-1.0...1.0] {
        var y x.line(m, c)
        
        ys.push(y)
    }
}
```