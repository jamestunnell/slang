# Arrays

An array is a fixed-type, dynamically-sized value sequence.

Array literals have their size and type determined from the initial expressions. The literal cannot be empty, since it is used to determine value type.

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

Lastly, the array can be both sized and initialized.
```
// [1, 2, ..., 11, 12]
var x array<int>(12, func(idx int) (val int) { val = idx+1 })
```