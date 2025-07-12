# Slang

Slang is a statically-typed scripting language.

## Fundamental Types

The fundamental data types are: `int`, `float`, `bool`, `string` (and `error`, `test`, and `time`?).

These data types are used to create constants, variables, structure fields, and function parameters.


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

### Tuples

Tuples store typed, ordered data in unnamed fields. A tuple is defined using the `tuple` keyword.

```
tuple StrFltInt(string float int)
```

A tuple value is created using the tuple name and a function call-like syntax where fields are specified by their order. Either all fields are be specified (non-empty tuple) or none (empty tuple). An empty tuple assigns zero values to each field.

```
var t1 StrFltInt("height" 65.7 18)  // non-empty tuple
var t2 StrFltInt                    // nil (empty) tuple
```

## Nil Values

Fundamental and compound data can be assigned values or left nil, and is equivalent to the zero value. A nil value can be used anywhere an assigned zero value can. For a compound type, a nil value means all constituent data is also nil (again, not invalid just equivalent to zero values).

## Functions

Functions execute code. A function is defined using the `func` keyword.

### Environment

A function environment starts with variables from input and output parameters. Output variables are assigned zero values.

Additional function-scope variables At each scope within the function body, new variables must be declared at the top of scope.

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