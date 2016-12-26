Dpackage main

import "fmt"

func main() {

	//create a slice of strings
	//containg a 	lenght and capacity of 5 elements
	slice1 := make([]string, 5)

	//create a slice of integer
	//contains a lenght of 3 and a capacity of 5 elements
	slice2 := make([]int, 3, 5)

	// create a slice of strings
	// contians a lenght and capacity of 5 elements
	slice3 := []string{"Red", "Yellow", "Black", "Green", "White"}

	// create a slice of integer, initialize the 100 elements with a value 44
	slice4 := []int{99: 44}

	// Remember : if you specify a value inside the [ ] operator, you're creating an array.
	// if you don't specify a value, you're creating a slice
	arrayD := [3]int{7, 9, 0} // creae an array of three integers
	sliceD := []int{7, 9, 0}  // create a slice of integers eith a lenght and capacity of 3

	// nil slice of strings
	var slice5 []string
	// create a empty slice of integer
	slice5 := make([]int, 0)

	// iterate over each elements and display each value
	for i, value := range sliceD {
		fmt.Printf("Index: %d  Value %d \n", i, value)

	}
	for i := 0; i < 5; i++ {
		fmt.Printf("array[%d] = %d \n", i, i)

	}

	fmt.Println("This is an integer of 5 elements!!")
	for i := 0; i < 5; i++ {

	}

}
