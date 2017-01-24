package exercise

import "fmt"

func greatest(v ...float64) float64 {
	max := -1000000.00
	for _, valore := range v {
		if valore > max {
			max = valore
		}

	}
	return max
}

func main() {

	lis := []float64{89.9, 454, 59495, 232, 455, 3445, 5452, 4209}
	fmt.Printf("max value is : %v", greatest(lis...))
}
