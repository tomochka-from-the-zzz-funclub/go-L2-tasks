package main

import "fmt"

type Command interface {
	Execute() string
}

type Light struct {
}

func (l *Light) TurnOn() string {
	return "Light is turned on"
}

func (l *Light) TurnOff() string {
	return "Light is turned off"
}

type TemperatureControl struct {
}

func (t *TemperatureControl) SetTemperature() string {
	return "Temperature is set"
}

type TurnOnLightCommand struct {
	light *Light
}

func (c *TurnOnLightCommand) Execute() string {
	return c.light.TurnOn()
}

type TurnOffLightCommand struct {
	light *Light
}

func (c *TurnOffLightCommand) Execute() string {
	return c.light.TurnOff()
}

type SetTemperatureCommand struct {
	tempControl *TemperatureControl
}

func (c *SetTemperatureCommand) Execute() string {
	return c.tempControl.SetTemperature()
}

func main() {
	// Создаем устройства
	light := &Light{}
	tempControl := &TemperatureControl{}

	// Создаем команды
	commands := []Command{
		&TurnOnLightCommand{light},
		&TurnOffLightCommand{light},
		&SetTemperatureCommand{tempControl},
	}

	// Выполняем команды
	for _, cmd := range commands {
		fmt.Println(cmd.Execute())
	}
}
