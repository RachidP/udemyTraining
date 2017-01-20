package main

import "fmt"

func main() {
	fmt.Println("before goroutines")

	go func() {

		for z := 0; z < 10; z++ {
			for i := 'A'; i < 'A'+26; i++ {
				fmt.Printf("%c", i)
			}
		}
	}()

	go func() {
		for z := 0; z < 10; z++ {
			for x := 'a'; x < 'a'+26; x++ {
				fmt.Printf("%c", x)

			}

		}

	}()
	fmt.Printf("after gorotimes \n")

}
