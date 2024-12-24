package main

import (
	"fmt"
)

type VehicleHandler interface {
	Handle(req VehicleRequest) string
	SetNext(h VehicleHandler)
}

type CarHandler struct {
	next VehicleHandler
}

func (c *CarHandler) SetNext(h VehicleHandler) {
	c.next = h
}

func (c *CarHandler) Handle(req VehicleRequest) string {
	if req.vehicleType == "car" {
		return fmt.Sprintf("Car: %s, Speed: %d km/h", req.name, req.speed)
	}
	if c.next == nil {
		return "No handler for this vehicle type."
	}
	return c.next.Handle(req)
}

type BikeHandler struct {
	next VehicleHandler
}

func (b *BikeHandler) SetNext(h VehicleHandler) {
	b.next = h
}

func (b *BikeHandler) Handle(req VehicleRequest) string {
	if req.vehicleType == "bike" {
		return fmt.Sprintf("Bike: %s, Speed: %d km/h", req.name, req.speed)
	}
	if b.next == nil {
		return "No handler for this vehicle type."
	}
	return b.next.Handle(req)
}

type TruckHandler struct {
	next VehicleHandler
}

func (t *TruckHandler) SetNext(h VehicleHandler) {
	t.next = h
}

func (t *TruckHandler) Handle(req VehicleRequest) string {
	if req.vehicleType == "truck" {
		return fmt.Sprintf("Truck: %s, Speed: %d km/h", req.name, req.speed)
	}
	if t.next == nil {
		return "No handler for this vehicle type."
	}
	return t.next.Handle(req)
}

type VehicleRequest struct {
	vehicleType string
	name        string
	speed       int
}

func main() {
	car := &CarHandler{}
	bike := &BikeHandler{}
	truck := &TruckHandler{}

	car.SetNext(bike)
	bike.SetNext(truck)

	fmt.Println(car.Handle(VehicleRequest{"car", "Toyota", 120}))
	fmt.Println(car.Handle(VehicleRequest{"bike", "Yamaha", 80}))
	fmt.Println(car.Handle(VehicleRequest{"truck", "Volvo", 90}))
	fmt.Println(car.Handle(VehicleRequest{"bus", "Mercedes", 100}))
}
