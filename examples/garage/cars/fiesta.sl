use Car from "../models"

func fiesta(
    color, vin str
    year int
    miles float) (car Car) {
    car = Car(year, make: "Ford", model: "Fiesta", color, vin, miles)
}
