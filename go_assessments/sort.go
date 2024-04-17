package main

import (
	"fmt"
)

func BubbleSort(numSlic []int) {

	for i := 0; i < 10-1; i++ {
		for j := 0; j < 10-i-1; j++ {
			if numSlic[j] > numSlic[j+1] {
				Swap(numSlic, j)
			}
		}

	}
	fmt.Println("Sorting numbers with Bubble sort is...")
	fmt.Println(numSlic)
}
func Swap(numSlic []int, index int) {
	temp := numSlic[index]
	numSlic[index] = numSlic[index+1]
	numSlic[index+1] = temp

}
func main() {
	var num int
	numSlic := make([]int, 10, 10)
	fmt.Println("Enter 10 Integer numbers")
	for i := 0; i < 10; i++ {
		fmt.Scan(&num)
		numSlic[i] = num
	}
	BubbleSort(numSlic)
}
