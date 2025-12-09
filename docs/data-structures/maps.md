# Maps (Associative Arrays)

A map is a fixed-type, dynamically sized associative array. They key type must have a Compare method.

An empty map must specify the type.

```
var m map<string int>

m.add("x" 120)
m.add("y" 121)
m.add("z" 122)

// 122
std.out << m.get("z")
```

The map key-value types can be inferred for non-empty literals. Key-value pairs are provided using tuples.

```
var m map[(12 6) (17 7)]

// 2
std.out << m.count()
```
