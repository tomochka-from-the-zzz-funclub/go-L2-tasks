package main

import "fmt"

type Strategy interface {
	CountCurrency(value float64) float64
}

type Client struct {
	Strategy
}

func (c *Client) SetStrategy(str string) {
	if str == "eur to rub" {
		c.Strategy = &EurToRub{}
	} else if str == "usd to rub" {
		c.Strategy = &UsdToRub{}
	}
}

type EurToRub struct {
}

func (etr *EurToRub) CountCurrency(num float64) float64 {
	return num * 62.73
}

type UsdToRub struct {
}

func (utr *UsdToRub) CountCurrency(num float64) float64 {
	return num * 61.12
}

func main() {
	client := Client{}
	client.SetStrategy("usd to rub")
	fmt.Println(client.CountCurrency(1000))
}
