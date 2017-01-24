package main

import "fmt"

func main() {

	for i := 0; i < 5500; i++ {
		fmt.Printf("%v - %v -  %v \n", i, string(i), []byte(string(i)))

	}

}
