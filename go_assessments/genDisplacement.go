package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	fn := func(t float64) float64 {

		return 0.5*(a*math.Pow(t, 2)) + ((v0 * t) + s0)
	}
	return fn

}
func main() {

	var a, v0, s0, t float64
	fmt.Println("Enter acceleration -> a, initial velocity ->v0, and initial displacement -> s0")
	fmt.Scan(&a, &v0, &s0)
	fmt.Println("Enter time ->t")
	fmt.Scan(&t)
	res := GenDisplaceFn(a, v0, s0)

	fmt.Println("Calculated result for s = ½ a t2 + vot + so: ", res(t))

}
