struct Garage(
    spots int
    cars array<Car>
)

func checkOut(
    g Garage
    c Car
) (ok bool) {
    ok = g.cars.delete(car)
}

func isFull (g Garage) (is bool) {
    is = g.cars.len() >= g.spots
}

func checkIn (
    g Garage
    c Car
) (ok bool) {
    if !g.full() {
        g.cars.add(c)

        ok = true
    }
}
