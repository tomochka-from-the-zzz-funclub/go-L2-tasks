package main

import (
	"fmt"
	"log"
)

// Product interface
type Vehicle interface {
	Name() string
}

// concrete product 1
type Car struct {
}

func NewCar() Vehicle {
	return &Car{}
}

func (c *Car) Name() string {
	return "car"
}

// concrete product 2
type Motorcycle struct {
}

func NewMotorcycle() Vehicle {
	return &Motorcycle{}
}

func (m *Motorcycle) Name() string {
	return "motorcycle"
}

// concrete product 3
type Truck struct {
}

func NewTruck() Vehicle {
	return &Truck{}
}

func (t *Truck) Name() string {
	return "truck"
}

// Creator interface
type VehicleCreator interface {
	CreateVehicle(str string) Vehicle
}

// ConcreteCreator struct
type ConcreteVehicleCreator struct {
}

func NewVehicleCreator() VehicleCreator {
	return &ConcreteVehicleCreator{}
}

func (c *ConcreteVehicleCreator) CreateVehicle(str string) Vehicle {
	var vehicle Vehicle
	switch str {
	case "car":
		vehicle = &Car{}
	case "motorcycle":
		vehicle = &Motorcycle{}
	case "truck":
		vehicle = &Truck{}
	default:
		log.Fatal("unknown vehicle type")
	}
	return vehicle
}

func main() {
	creator := NewVehicleCreator()
	fmt.Println(creator.CreateVehicle("car").Name())
	fmt.Println(creator.CreateVehicle("motorcycle").Name())
	fmt.Println(creator.CreateVehicle("truck").Name())
	// Uncommenting the next line will cause the program to log an error
	// fmt.Println(creator.CreateVehicle("bicycle").Name())
}
