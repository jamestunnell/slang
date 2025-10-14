use Car from "../models"

func mustang(
    color, vin str
    year int
    miles flt) (car Car) {
    car = Car(year, make: "Ford", model: "Mustang", color, vin, miles)
}
