use "date"

class Car {
  field Make string
  field Model string
  field Year int

  field miles float
  field nextOilChange float

  method init() {
    this.miles = 0.0
    this.nextOilChange = 5000.0
  }

  method Age() int {
    return date.Today().Year - this.Year
  }

  method String() string {
    return "${this.Color} ${this.Year} ${this.Make} ${this.Model}"
  }

  method Drive(miles float) {
    this.miles = this.miles + miles
  }

  method NeedsOilChange() bool {
    return this.miles >= this.nextOilChange
  }

  method ChangeOil() {
    this.nextOilChange = this.miles + 5000.0
  }
}

func TestCarAge(t test.Test) {
  car1 = Car.New(Make: "Buick", Model: "LeSabre", Year: 2000)
  car2 = Car.New(Make: "Cadillac", Model: "Escalade", Year: 2003)

  t.AssertGreater(car2.Age(), car1.Age())
}

func TestCarOilLife(t test.Test) {
  car = Car.New(Make: "Ford", Model: "Bronco", Year: 1990)
  
  t.AssertFalse(car.NeedsOilChange())

  car.Drive(1000)

  t.AssertFalse(car.NeedsOilChange())

  car.Drive(3999.9)

  t.AssertFalse(car.NeedsOilChange())

  car.Drive(1)

  t.AssertTrue(car.NeedsOilChange())

  car.ChangeOil()

  t.AssertFalse(car.NeedsOilChange())
}
