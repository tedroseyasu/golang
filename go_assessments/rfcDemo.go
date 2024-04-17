package main

import (
	"fmt"
	"io/ioutil"
	//"net/http"
)

func main() {

	//http.Get(www.google.com)
	t1 := []byte("This is for test")

	err := ioutil.WriteFile("test.txt", t1, 0777)

	t, err := ioutil.ReadFile("test.txt")
	fmt.Println(string(t))

	if err != nil {

		fmt.Println("File opening Error!")
	}
}
