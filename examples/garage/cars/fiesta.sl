use "../models"

const make "Ford"

func NewFiesta(
    vin str
    year int
    color str) (car z.Car) {
    car = z.Car(
        Make make
        Model "Fiesta"
        Color color
        Year year
        Miles 0.0
        VIN vin
    )
}
