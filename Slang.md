# Slang

Slang is a statically-typed scripting language.

### Built-in Types

Built-in data types are: `int`, `float`, `bool`, `string` (and `time`?).

These data types are used to create constants, variables, structure fields, and function parameters.

### Structures

Structures store ordered data in named fields. A structure is defined using the `struct` keyword.

```
struct Point2D(
    x,y float
)
```

A structure value is created using the struct name and a function call-like syntax where fields are specified with keyword arguments. Either all fields are be specified (non-empty struct) or none are (empty struct). An empty struct assigns zero values to each field.

```
var p1 Point2D(x: 10, y: 22.8)  // non-empty struct
var p2 Point2D()                // empty struct
```

### Tuples

Tuples store ordered data in unnamed fields. A tuple is defined using the `tuple` keyword.

```
tuple NameVal(
    string
    float
)
```

A tuple value is created using the tuple name and a function call-like syntax where fields are specified by their order. Either all fields are be specified (non-empty tuple) or none (empty tuple). An empty tuple assigns zero values to each field.

```
var nv1 NameVal("height" 65.7)  // non-empty tuple
var nv2 NameVal()               // empty tuple
```

## Functions

Functions execute code. A function is defined using the `func` keyword.

### Environment

A function environment starts with variables from input and output parameters. Output variables are assigned zero values.

Additional function-scope variables At each scope within the function body, new variables must be declared at the top of scope.

### Signature

Function signature is the union of the sets of input and output parameter types. If either input types or output types are different between two functions, then they will have different signatures. The order of types does not matter.

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
