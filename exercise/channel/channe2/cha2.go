package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("before goroutines")
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for z := 0; z < 3; z++ {
			for i := 'A'; i < 'A'+26; i++ {
				fmt.Printf("%c ", i)
			}
		}

	}()

	go func() {
		defer wg.Done()
		for z := 0; z < 3; z++ {
			for x := 'a'; x < 'a'+26; x++ {
				fmt.Printf("%c ", x)

			}

		}
	}()

	fmt.Printf("\n after gorotimes  and before wait\n")
	wg.Wait()
	fmt.Println(" \nafter wait")

}
