// Keeps a running total.
// Updates the total after each operation.
struct Calculator(total flt)

func add(
    c Calculator
    x flt) {
    c.total = c.total + x
}

func sub(
    c Calculator
    x flt) {
    c.total = c.total - x
}

func mul(
    c Calculator
    x flt) {
    c.total = c.total * x
}

func div(
    c Calculator
    x flt) {
    c.total = c.total / x
}

func clear(c Calculator) {
    c.total = 0.0
}

func total(c Calculator) (result flt) {
    result = c.total
}

func testCalculator1(t test.T) {
    var c Calculator(total: 0.0)
    
    c.mul(2.0)
    c.add(1.5)

    t.Approx(
        actual c.total()
        expected 1.5)
}

func testCalculator2(t test.T) {
    c = Calculator(total: 1.1)

    c.mul(22.2)
    c.sub(7.6)
    c.div(10)
    c.add(0.5)

    t.Approx(
        actual: c.total()
        expected: 23.94)
}