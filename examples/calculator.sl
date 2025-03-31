// Keeps a running total.
// Updates the total after each operation.
struct Calculator {
    total float
}

func Add(c Calculator, x float) {
    c.total = c.total + x
}

func Sub(c Calculator, x float) {
    c.total = c.total - x
}

func Mul(c Calculator, x float) {
    c.total = c.total * x
}

func Div(c Calculator, x float) {
    c.total = c.total / x
}

func Clear(c Calculator) {
    c.total = 0.0
}

func Total(c Calculator) float {
    return c.total
}