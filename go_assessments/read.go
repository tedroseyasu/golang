package main

import (
	"bufio"
	"strings"

	//"encoding/json"
	"fmt"
	"os"
)

type Name struct {
	fname string
	lname string
}

func main() {

	var filename string
	var readName Name
	nameOnFile := make([]Name, 0, 20)

	fmt.Println("Enter the file name: ")
	fmt.Scan(&filename)
	openFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("FILE READING ERROR.......", err)

	}
	readLines := bufio.NewScanner(openFile)

	for readLines.Scan() {
		readWord := strings.Split(readLines.Text(), " ")
		readName.fname = readWord[0]
		readName.lname = readWord[1]
		nameOnFile = append(nameOnFile, readName)
	}
	fmt.Println("--------------------")
	for _, val := range nameOnFile {

		fmt.Println("First Name: ", val.fname)
		fmt.Println("Last Name: ", val.lname)

	}
	fmt.Println("--------------------")

}
