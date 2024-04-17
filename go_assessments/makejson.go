package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	
)

func main() {

	var dat map[string]string
	dat = make(map[string]string)
	var name string

	fmt.Println("Enter your First Name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter your Address: ")
	reader := bufio.NewReader(os.Stdin)
	address, err := reader.ReadString('\n')
	address = strings.TrimSpace(address)
	if err != nil {
		log.Fatal(err)
	}

	dat["name"] = name
	dat["address"] = address

	jval, err2 := json.MarshalIndent(dat, "", "\t")

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("JSON version of the Map values: ")
	fmt.Println(string(jval))

}
