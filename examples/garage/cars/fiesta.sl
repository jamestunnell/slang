use "../models"

func NewFiesta(vin string, year int, color string) (car models.Car) {
    car = models.Car(
        Make: "Ford",
        Model: "Fiesta",
        Color: color,
        Year: year,
        Miles: 0.0,
        VIN: vin
    )
}
