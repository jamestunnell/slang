# Structures

Structures store ordered data in named fields. A structure is defined using the `struct` keyword.

```
struct Point2D(
    x, y float
)
```

A structure value is created using the struct name and a function call-like syntax where fields are specified with keyword arguments. Either all fields are be specified (non-empty struct) or none are (nil, empty struct). An empty struct assigns zero values to each field.

```
var p1 Point2D(x:10 y:22.8) // non-empty, fields by keyword args
var p1 Point2D(10 22.8)     // non-empty, fields by ordinal args
var p2 Point2D              // nil (empty) struct
```

## Anonymous Structs

A struct value can be created without specifying a defined type. An anonymous type will be defined based on the fields.

```
var x (x:10 y:22.8)  // anonymous type is struct(x, y float)
```

