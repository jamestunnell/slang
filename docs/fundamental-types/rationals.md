# Rational Values

`rat` is a rational number, that is a quotient of two integers.

```
var q rat(20 3)

// 20/3
std.out << q

// 6.6666666666666666666666666666667
std.out << q.toFloat()

// 13/3
std.out << q.sub(rat(7 3))
```