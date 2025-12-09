# Tuples

Tuples store positional data (fields are typed and ordered, but not named). A tuple type is defined using the `tuple` keyword.

```
tuple Point2D(float float)
```

A tuple value is created using the tuple name and a function call-like syntax where fields are specified by their order. Either all fields are be specified (non-empty tuple) or none (empty tuple). An empty tuple assigns zero values to each field.

```
var t1 Point2D(65.7 29.9)  // non-empty tuple
var t2 Point2D             // nil (empty) tuple
```

## Anonymous Tuples

A tuple value can be created without specifying a defined type. An anonymous type will be defined based on the fields.

```
var x (67.5 "okay")  // anonymous type is tuple(float string)
```