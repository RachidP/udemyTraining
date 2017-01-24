package main

import "fmt"

func fib(num int) []int {

	slice := make([]int, num)
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		if i < 2 {
			slice[i] = 1

		} else {

			slice[i] = slice[i-1] + slice[i-2]
		}

	}

	return slice
}

func main() {

	fmt.Printf("value is : %v", fib(9))
}
