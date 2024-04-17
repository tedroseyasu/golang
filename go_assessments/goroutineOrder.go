package main

import (
	"fmt"
	"sort"
)

func SortFunc(arr []int, channel chan []int) {
	sort.Ints(arr)
	channel <- arr
}

func splitFunc(arr []int) (arr1 []int, arr2 []int) {

	splitSize := len(arr) / 2
	remainSize := len(arr) - splitSize

	arr1 = make([]int, remainSize)
	arr2 = make([]int, splitSize)
	copy(arr1, arr[splitSize:])
	copy(arr2, arr[:splitSize])
	return arr1, arr2
}
func sortVal(arr []int, c chan []int) {
	sort.Ints(arr)
	c <- arr
}
func main() {

	var le int
	fmt.Println("Enter the size of the inputs:")
	fmt.Scan(&le)

	arr := make([]int, le)

	fmt.Println("Enter the values:")
	for input := 0; input < le; input++ {
		fmt.Scan(&arr[input])
	}

	firstHalf, secondHalf := splitFunc(arr)
	firstHalf1, firstHalf2 := splitFunc(firstHalf)
	secondHalf1, secondHalf2 := splitFunc(secondHalf)

	c := make(chan []int)

	//4  goroutine
	go sortVal(firstHalf1, c)
	go sortVal(firstHalf2, c)
	go sortVal(secondHalf1, c)
	go sortVal(secondHalf2, c)

	sorted1 := <-c
	sorted2 := <-c
	sorted3 := <-c
	sorted4 := <-c

	fmt.Println("Sorted partition1: ", sorted1)
	fmt.Println("Sorted partition2: ", sorted2)
	fmt.Println("Sorted partition3: ", sorted3)
	fmt.Println("Sorted partition4: ", sorted4)

	merge1 := append(sorted1, sorted2...)
	merge2 := append(sorted3, sorted4...)
	lastMerge := append(merge1, merge2...)

	sort.Ints(lastMerge)
	fmt.Println("After running 4 goroutine the values are: ", lastMerge)

}
