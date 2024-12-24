package main

import "fmt"

type Visitor interface {
	DriveCar(c *Car) string
	RideMotorcycle(m *Motorcycle) string
}

type Vehicle interface {
	Accept(v Visitor) string
}

type ConcreteVisitor struct {
}

func (c *ConcreteVisitor) DriveCar(car *Car) string {
	return car.Drive()
}

func (c *ConcreteVisitor) RideMotorcycle(motorcycle *Motorcycle) string {
	return motorcycle.Ride()
}

type Car struct {
}

func (c *Car) Drive() string {
	return "Driving a car"
}

func (c *Car) Accept(v Visitor) string {
	return v.DriveCar(c)
}

type Motorcycle struct {
}

func (m *Motorcycle) Ride() string {
	return "Riding a motorcycle"
}

func (m *Motorcycle) Accept(v Visitor) string {
	return v.RideMotorcycle(m)
}

func main() {
	// make queue of objects for visitor
	vehicles := []Vehicle{&Car{}, &Motorcycle{}}
	for _, v := range vehicles {
		cv := ConcreteVisitor{}
		fmt.Println(v.Accept(&cv))
	}
}
