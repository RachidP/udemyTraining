package main

import "fmt"

func main() {

	for index := range fact(10) {
		fmt.Println(index)
	}

}

func fact(num int) chan int {
	c := make(chan int)

	go func() {
		v := 1
		if num > 1 {
			for i := 1; i <= num; i++ {
				v *= i
			}
			c <- v
			close(c)
		}
	}()

	return c
}
