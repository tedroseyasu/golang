package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	//"unicode"
)

func main() {

	var num string
	var slic []int

	slic = make([]int, 0, 3)
	for {

		fmt.Println("Enter num: or X to Exit ")
		fmt.Scanln(&num)
		v, err := strconv.Atoi(num)

		if err != nil {
			fmt.Println("")
		}
		if num == "X" {
			fmt.Println("Thank you for using the program!")
			break
		} else {
			isnum := regexp.MustCompile(`\d`).MatchString(num)

			if isnum == true {
				slic = append(slic, v)
			}

			sort.Ints(slic)
			fmt.Println("The value you entered is: ", slic)
		}

	}

}
