package main

type Director struct {
	b Builder
}

func NewDirector() *Director {
	return &Director{b: &ConcreteBuilder{c: new(Car)}}
}

func (d *Director) Build(make, model, color string, year int) *Car {
	d.b.SetMake(make)
	d.b.SetModel(model)
	d.b.SetColor(color)
	d.b.SetYear(year)
	return d.b.GetCar()
}

type Builder interface {
	SetMake(make string)
	SetModel(model string)
	SetColor(color string)
	SetYear(year int)
	GetCar() *Car
}

type ConcreteBuilder struct {
	c *Car
}

func (c *ConcreteBuilder) SetMake(make string) {
	c.c.Make = make
}

func (c *ConcreteBuilder) SetModel(model string) {
	c.c.Model = model
}

func (c *ConcreteBuilder) SetColor(color string) {
	c.c.Color = color
}

func (c *ConcreteBuilder) SetYear(year int) {
	c.c.Year = year
}

func (c *ConcreteBuilder) GetCar() *Car {
	return c.c
}

type Car struct {
	Make, Model, Color string
	Year               int
}

// func main() {
// 	d := NewDirector()
// 	car := d.Build("Toyota", "Camry", "Red", 2021)
// 	fmt.Printf("Car: %s %s, Color: %s, Year: %d\n", car.Make, car.Model, car.Color, car.Year)
// }
