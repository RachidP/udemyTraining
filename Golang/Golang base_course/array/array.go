package main

import "fmt"

func main() {

	//declare an array of five elements
	var array [5]int
	//declare an intger array of 5 elements
	array2 := [5]int{10, 20, 30, 40, 50}

	array3 := []int{7, 9, 9, 33, 86}

	array4 := [5]int{2: 44, 4: 44}

	for i := 0; i < len(array); i++ {
		fmt.Printf("array[%d] = %d \n", i, array[i])

	}

	fmt.Println("This is an integer of 5 elements!!")
	for i := 0; i < len(array2); i++ {
		fmt.Printf("array2[%d]  = %d \n", i, array2[i])
		fmt.Printf("array3[%d]  = %d \n", i, array3[i])
		fmt.Printf("array4[%d]  = %d \n", i, array4[i])

	}

	// declare an integer pointer array of 5 elements
	array5 := [5]*int{0: new(int), 4: new(int)}
	// assign value to index 0 and 1
	*array5[0] = 99
	*array5[4] = 100

	fmt.Println("This is an integer pointer array of 5 elements!!")
	for i := 0; i < len(array5); i++ {
		fmt.Printf("array5[%d]  = %d \n", i, array5[i])

	}

	fmt.Println("Array of strings!!")
	arr6 := [5]string{"Red", "Yellow", "Black", "Green", "White"}
	var arr7 [5]string
	arr7 = arr6
	for i := 0; i < len(arr6); i++ {
		fmt.Printf("array6[%d] =%s \n", i, arr6[i])
		fmt.Printf("array7[%d] =%s \n", i, arr7[i])

	}

}
