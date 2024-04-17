package main

import (
	"fmt"
)

type Animal interface {
	Name() string
	Eat()
	Speak()
	Move()
}
type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}
type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}
type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println("Cow eats: ", c.food)
}
func (c Cow) Move() {
	fmt.Println("Cow moves: ", c.locomotion)
}
func (c Cow) Speak() {
	fmt.Println("Cow speak: ", c.noise)
}
func (b Bird) Eat() {
	fmt.Println("Birds eats: ", b.food)
}
func (b Bird) Move() {
	fmt.Println("Birds moves: ", b.locomotion)
}
func (b Bird) Speak() {
	fmt.Println("Bird speak: ", b.noise)
}
func (s Snake) Eat() {
	fmt.Println("Snake eats: ", s.food)
}
func (s Snake) Move() {
	fmt.Println("Snake moves: ", s.locomotion)
}
func (s Snake) Speak() {
	fmt.Println("Snake speak: ", s.noise)
}
func (c Cow) Name() string {
	return c.name
}

func (b Bird) Name() string {
	return b.name
}

func (s Snake) Name() string {
	return s.name
}

func main() {

	animalCol := make(map[string]Animal, 0)

	var newOrQuery, name, typeOf string
	for {
		fmt.Println("Enter 'newanimal', the animal 'name' and type of animal('cow,bird,snake')")
		fmt.Println("OR Enter 'query' and 'name' of the animal and Methods of animal('eat, move, speak')")
		fmt.Print(">")
		fmt.Scan(&newOrQuery, &name, &typeOf)

		switch newOrQuery {

		case "newanimal":
			switch typeOf {
			case "cow":
				animalCol[name] = Cow{name, "grass", "walk", "moo"}
			case "bird":
				animalCol[name] = Bird{name, "worms", "fly", "peep"}
			case "snake":
				animalCol[name] = Snake{name, "mice", "slither", "hsss"}
			}
			fmt.Println("Created it!")

		case "query":
			animal := animalCol[name]
			switch typeOf {
			case "eat":
				animal.Eat()
			case "speak":
				animal.Speak()
			case "move":
				animal.Move()
			}
		default:
			fmt.Println("Please enter newanimal/query")
		}

	}
}
