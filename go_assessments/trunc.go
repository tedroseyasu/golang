package main

import (
	"fmt"
)

func main() {

	var flt_Val float32
	var res int
	fmt.Println("Enter the Float value: ")
	//getting value from console
	fmt.Scan(&flt_Val)

	res = int(flt_Val)
	//printing integer value
	fmt.Println("the integer value is: ", res)

}
