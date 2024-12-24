package main

type Farm struct {
	cow     *Cow
	sheep   *Sheep
	chicken *Chicken
}

func NewFarm() *Farm {
	return &Farm{
		cow:     &Cow{},
		sheep:   &Sheep{},
		chicken: &Chicken{},
	}
}

type Cow struct {
}

func (c *Cow) MakeSound() string {
	return "Moo"
}

type Sheep struct {
}

func (s *Sheep) MakeSound() string {
	return "Baa"
}

type Chicken struct {
}

func (ch *Chicken) MakeSound() string {
	return "Cluck"
}

// func main() {
// 	farm := NewFarm()
// 	fmt.Println(farm.cow.MakeSound())
// 	fmt.Println(farm.sheep.MakeSound())
// 	fmt.Println(farm.chicken.MakeSound())
// }
