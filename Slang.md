# Slang

Slang is a statically-typed scripting language.

## Fundamental Types

The fundamental data types are: `int`, `float`, `bool`, `string`, `cmp` (and `error`, `test`, and `time`?).

These data types are used to create constants, variables, structure fields, array/tuple elements, and function parameters.

### string

Strings are immutable.

String interpolation occurs when an expression is placed in double braces, like `var str "my name is {{getName()}}"`. The expression value must be string or have a `toString`  method.

### error

`error` is a built-in type used to convey failure. It will either have nil value, if unassigned, or it will contain an error category and message.

```
use "errors"

func verifyPositive(x int) (err error) {
    if x <= 0 {
        err = errors.constraintViolation("{{x}} is not positive")
    }
}
```

### cmp

`cmp` is a built-in type is used to compare ordinal values. It will either have nil value, if unassigned, or one of the three valid values: `lt`, `eq`, and `gt`.

## Arrays

An array is a fixed-type, dynamically-sized value sequence.

Array literals have their size and type determined from the initial expressions. The literal cannot be empty, since it is used to determine the value type.

```
const Z 13
var x {6,7,8*Z}
```

Non-literal arrays can be created in a few ways. They can be left empty:
```
var x array<int>
```

Or they can be sized but with values left unassigned.
```
var x array<int>(12)
```

Lastly, the array can be both sized and initialized, using an init func.
```
var x array<int>(12, func(idx int) (val int) { val = idx + 1 })
```

## Compound Types

A compound type contains one or more fundamental types.

### Structures

Structures store ordered data in named fields. A structure is defined using the `struct` keyword.

```
struct Point2D(
    x, y float
)
```

A structure value is created using the struct name and a function call-like syntax where fields are specified with keyword arguments. Either all fields are be specified (non-empty struct) or none are (nil, empty struct). An empty struct assigns zero values to each field.

```
var p1 Point2D(x:10 y:22.8)  // non-empty struct
var p2 Point2D               // nil (empty) struct
```

#### Anonymous Structs

A struct value can be created without specifying a defined type. An anonymous type will be defined based on the fields.

```
var x (x:10 y:22.8)  // anonymous type is struct(x, y float)
```


### Tuples

Tuples store positional data (fields are typed and ordered, but not named). A tuple type is defined using the `tuple` keyword.

```
tuple Point2D(float float)
```

A tuple value is created using the tuple name and a function call-like syntax where fields are specified by their order. Either all fields are be specified (non-empty tuple) or none (empty tuple). An empty tuple assigns zero values to each field.

```
var t1 Point2D(65.7 29.9)  // non-empty tuple
var t2 Point2D             // nil (empty) tuple
```

#### Anonymous Tuples

A tuple value can be created without specifying a defined type. An anonymous type will be defined based on the fields.

```
var x (67.5 "okay")  // anonymous type is tuple(float string)
```

#### Tables

Data can be organized into tables. Good for test and seed data, but also for any data coming from a database, or for data that can benefit from more organization: A column is array-like, but includes the column name. A row is struct-like, but ???

Example:
```
struct Poster(
    franchise, variant string
    price decimal
)

var posters table<Poster>(
    "Planet of the Apes"    "Statue of Liberty"         11.99
    "Planet of the Apes"    "Marcus, Head of Security"  8.99
    "One Piece"             "Luffy's Bounty"            12.99
    "One Piece"             "Cross Guild"               10.99
    "Studio Ghibli"         "Characters Collage"        15.99
    "Studio Ghibli"         "Kiki's Delivery Service"   12.99
)

func findPostersWithPriceInRange(min, max decimal) (found rows<Poster>) {
    found = posters.Where(func(p Poster) (matches bool) {
        matches = p.price >= min and p.price <= max
    }
}

func allPosterPrices() (prices column<decimal>) {
    prices = posters.prices
}

```

## Nil Values

Fundamental and compound data can be assigned values or left nil, and is equivalent to the zero value. A nil value can be used anywhere an assigned zero value can. For a compound type, a nil value means all constituent data is also nil (again, not invalid just equivalent to zero values).

## Functions

Functions execute code. A function is defined using the `func` keyword.

### Environment

A function environment starts with variables from input and output parameters. Output variables are assigned zero values if they are not assigned values at the top of function scope.

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

Function signature is the sequence of input and output parameter types. If the sequence of either input types or output types are different between two functions, then they will have different signatures. The order of types does not matter.

### Parameters

A function parameter has name and type. 

Within a set of parameters (input or output) types must be unique. 
This constraint has several effects:
* function inputs can be passed in any order
* function outputs can be collected in any order
* any function can be used as a method with any of the arguments
* any function can be used as a method with any tuple of the arguments

### Overloading

Functions can have the same name as long as the signature is different.

### Return Statement

Function execution can be terminated early using the `return` keyword. Output parameter values must be set (or left at their zero values) before returning.

### Calling Convention

Some examples that demonstrate the calling conventions, using the function below.

```
func linearEqn(x, slope, intercept int) (y int) {
    y = x * slope + intercept
}
```

#### Ordinal Arguments

```
func main() {
    const m 2.5
    const c -0.7
    var ys ary<float>

    for x in [-1.0...1.0] {
        var y linearEqn(x, m, c)
        
        ys.push(y)
    }
}
```

#### Auto-Keyword Arguments

```
func main() {
    const slope 2.5
    const intercept -0.7
    var ys ary<float>

    for x in [-1.0...1.0] {
        var y linearEqn((x, m, c))
        
        ys.push(y)
    }
}
```

#### Mixed-Keyword Arguments

```
func main() {
    const m 2.5
    const c -0.7
    var ys ary<float>

    for x in [-1.0...1.0] {
        var y linearEqn((x, slope: m, intercept: c))
        
        ys.push(y)
    }
}
```

#### Method Call Syntax

```
func main() {
    const m 2.5
    const c -0.7
    var ys ary<float>

    for x in [-1.0...1.0] {
        var y x.linearEqn(m, c)
        
        ys.push(y)
    }
}
```

## Blocks and `yield` statements

A block is a lexical scope that is tied to an invocation of a coroutine. As the coroutine executes, it can yield values to the block as it is executed, like a function call. Block input parameters are automatically determined from the declared yield parameters of the coroutine. These auto-input params are accessed using the `$` placeholder. 

For example:
```
use "fmt"

struct Poster(
    franchise, variant string
    price decimal
)

func describe(p Poster) returns (out string) {
    const price fmt.currency(p.price)
    out = "Don't miss out on '{{p.franchise}}: {{p.variant}}' for only ${{$.p.price}}!"
}

func cheapest(
    posters ary<Poster>
    n pint
) returns (cheapestPosters iter<Poster>) {
    cheapestPosters = posters.sortBy($.price).take(n)
}

func allPosters() returns (posters ary<Poster>) {
    posters = [
        Poster("Planet of the Apes", "Statue of Liberty", 11.99)
        Poster("Planet of the Apes", "Marcus, Head of Security", 8.99)
        Poster("One Piece", "Luffy's Bounty", 12.99)
        Poster("One Piece", "Cross Guild", 10.99)
        Poster("Studio Ghibli", "Characters Collage", 15.99)
        Poster("Studio Ghibli", "Kiki's Delivery Service", 12.99)
    ]
}

func main() {
    allPosters().cheapest(3).each() {
        cout << $.describe()
    }
}
```

###