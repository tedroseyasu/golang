package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Println("Enter a String value")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	str1 := strings.TrimSpace(str)

	if (strings.HasPrefix(strings.ToUpper(str1), strings.ToUpper("i")) && strings.HasSuffix(strings.ToUpper(str1), strings.ToUpper("n"))) && (strings.Contains(strings.ToUpper(str1), strings.ToUpper("a"))) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
	fmt.Println("You Entered:", str1)

}
