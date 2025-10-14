struct Car(
    year int
    make, model, color, vin str
    miles flt)

func drive(
    c Car
    miles flt) {
    c.miles = c.miles + miles
}

func compareVIN(a, b Car) (result cmp) {
    result = a.vin <=> b.vin
}