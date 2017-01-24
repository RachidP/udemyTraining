package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Printf("number of physical processors %v \n", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	ch := make(chan rune)
	done := make(chan bool)

	go func() {
		for z := 0; z < 2; z++ {
			for i := 'A'; i < 'A'+26; i++ {
				ch <- i
			}
		}
		done <- true
	}()

	go func() {

		for z := 0; z < 2; z++ {
			for x := 'a'; x < 'a'+26; x++ {
				ch <- x

			}
		}
		done <- true
	}()

	go func() {
		<-done
		<-done

		close(ch)

	}()

	fmt.Printf("\n after gorotimes  and before wait\n")

	for val := range ch {
		temp := val
		fmt.Printf("%c  ", temp)
	}

	fmt.Println(" \nafter wait")

}
