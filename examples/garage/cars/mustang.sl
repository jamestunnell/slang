use "../models"

func NewMustang(
    vin str
    year int
    color str) (car models.Car) {
    car = models.Car(
        Make: make
        Model: "Mustang"
        Color: color
        Year: year
        Miles: 0.0
        VIN: vin
    )
}
