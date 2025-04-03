use "../models"

func NewMustang(vin string, year int, color string) (car models.Car) {
    car = models.Car(
        Make: "Ford",
        Model: "Mustang",
        Color: color,
        Year: year,
        Miles: 0.0,
        VIN: vin
    )
}
