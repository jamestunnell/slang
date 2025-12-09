# Sets

A set is a fixed-type, dynamically sized value set.

An empty map must specify the type.
```
var s set<float>

s.add(2.5)
s.add(2.5)
s.add(2.6)

// set[2.5 2.6]
std.out << s
```

The set element type can be inferred for non-empty literals.

```
var s set[12.3 12.4]

// 2
std.out << s.count()
```