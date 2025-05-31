struct Garage (
    Cars ary<Car>
)

func Leave(
    g Garage
    c Car) {
    g.Cars.Delete(c)
}

func Return(
    g Garage
    c Car) {
    g.Cars.Add(c)
}