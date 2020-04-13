package main

import (
	"fmt"
)

// Animal ...
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat Method
func (animal *Animal) Eat() string {
	return animal.food
}

// Move Method
func (animal *Animal) Move() string {
	return animal.locomotion
}

// Speak Method
func (animal *Animal) Speak() string {
	return animal.noise
}

func main() {
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}

	mapAnimalToStruct := make(map[string]*Animal)
	mapAnimalToStruct["snake"] = &snake
	mapAnimalToStruct["cow"] = &cow
	mapAnimalToStruct["bird"] = &bird

	for {
		var animal, method string
		fmt.Print(">")
		fmt.Scan(&animal, &method)
		switch method {
		case "eat":
			fmt.Println(mapAnimalToStruct[animal].Eat())
			break
		case "move":
			fmt.Println(mapAnimalToStruct[animal].Move())
			break
		case "speak":
			fmt.Println(mapAnimalToStruct[animal].Speak())
			break
		}

	}

}
