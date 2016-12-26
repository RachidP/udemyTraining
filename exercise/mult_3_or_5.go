package main

import "fmt"

func multiplicazione() int {

	var max, v3, v5 int

	for i := 1; i*3 < 1000; i++ {
		v3 = i * 3
		v5 = i * 5
		if !(v3%5 == 0) {
			max += v3
		}
		if v5 < 1000 {
			max += v5
		}
	}

	return max
}

func main() {

	fmt.Printf("value is : %v", multiplicazione())
}
