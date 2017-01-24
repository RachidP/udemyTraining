package exercise

import "fmt"

func calc(val int) (int, bool) {

	var n bool
	if val%2 == 0 {
		n = true
	} else {
		n = false
	}
	return val / 2, n
}

func main() {
	var even bool
	var res int
	num := 6
	res, even = calc(num)
	fmt.Printf("number = %d, the %d/2=%d   is it even?%v", num, num, res, even)
}
