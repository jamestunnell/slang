use "../models"

func NewFiesta(vin str, year int, color str) (car models.Car) {
    car = models.Car(
        Make: "Ford",
        Model: "Fiesta",
        Color: color,
        Year: year,
        Miles: 0.0,
        VIN: vin
    )
}
