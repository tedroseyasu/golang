package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (e Animal) Eat() {
	fmt.Println(e.food)
}

func (m Animal) Move() {
	fmt.Println(m.locomotion)
}

func (s Animal) Speak() {
	fmt.Println(s.noise)
}
func main() {

	var Animal_Col = map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"}}
	var typeAnimal string
	var ems string
	for {
		fmt.Println("Enter the Animal and your request [eat, move, speak] >")
		fmt.Scan(&typeAnimal, &ems)
		val := Animal_Col[typeAnimal]

		if ems == "eat" {
			val.Eat()
		}
		if ems == "move" {
			val.Move()
		}
		if ems == "speak" {
			val.Speak()
		}
	}

}
