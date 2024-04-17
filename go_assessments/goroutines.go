package main

import (
	"fmt"
	"time"
)

// global variable that set to 5
var x int = 5

// Function 1
func func1() {
	x = x + 2
	//fmt.Println("Value for x is: ", x)
}

// Function 2
func func2() {
	x = x + 1
	fmt.Println("Value for x is: ", x)
}

// Main function
func main() {

	//Go Rountines [Creating thread]
	/*Race conditions occur when goroutines simultaneously read and write to the same
	shared memory space without synchronization mechanisms and unexpected result.
	e.g.: value of x is initially 5, then after executing the two function the result is still 5, after the two
	function end of the execution it changes the value.

	*/
	go func1()
	go func2()

	//enven-though the value is changed through func1 and func2 the value still 5
	fmt.Println("After the two goroutine function, the value of x is: ", x)

	time.Sleep(time.Second)

	fmt.Println("After the two goroutine function finished, the value of x2 is: ", x)

}
