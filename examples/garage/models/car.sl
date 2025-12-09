struct Car(
    year int
    make, model, color, vin str
    miles float)

func drive(
    c Car
    miles float) {
    c.miles = c.miles + miles
}

func compareVIN(a, b Car) (result cmp) {
    result = a.vin <=> b.vin
}