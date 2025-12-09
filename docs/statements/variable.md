# Variables

A variable is a value that can be changed at runtime. Variables are defined using the `var` keyword.

Initialization can be an expression (type inferred) or just the variable type (zero value will be used).

```
func sum(vals array<int>) int {
    var x int // auto-initialize to 0

    vals.eachValue | x = x.add($)
}
```

A variable must be declared in the top-of-scope area, after any constants and nested functions.

```
// mul4add2 modifies each value (in place) by multiplying by 4 then adding 2.
func mul4add2(vals array<float>) {
    func logEvery100th(
        in, out float
        index int) {
        if idx.mod(100).eq(0) {
            std.out << "[${index}]: in=${in} out=${out}"
        }
    }

    var in float
    var out float

    vals.each {
        in = $.value
        out = $.value.mul(4.0).add(2.0)

        logEvery100th(in: in, out: out, index: $.index)
        
        vals[$.index] = out
    }
}
```
