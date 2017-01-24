package main

import (
	"fmt"
	"sort"
)

func main() {
	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sR := []string{"Zeno", "John", "Al", "Jenny"}

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	nR := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	orderS(s)
	revOrderS(sR)
	//fmt.Printf("type people= %T  \n", studyGroup)
	orderP(studyGroup)
	orderInt(n)
	revOrderInt(nR)
}

func orderS(s []string) {
	fmt.Println("----------------------orderS------------------------------------------")

	fmt.Printf("Slice s before sorting:---> %v \n", s)
	sort.Strings(s)
	fmt.Printf("Slice s After sorting: --->%v \n", s)

}

func revOrderS(s []string) {
	fmt.Println("------------------revOrderS----------------------------------------------")

	fmt.Printf("Slice s before sorting:---> %v \n", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Printf("Slice s After sorting: --->%v \n", s)

}

func orderP(s []string) {
	fmt.Println("----------------------orderP------------------------------------------")
	fmt.Printf("s before sorting:---> %v \n", s)
	sort.Strings(s)
	fmt.Printf("s After sorting: --->%v \n", s)

}
func orderInt(s []int) {
	fmt.Println("--------------------orderInt--------------------------------------------")
	fmt.Printf(" s before sorting:---> %v \n", s)
	sort.Ints(s)
	fmt.Printf(" s After sorting: --->%v \n", s)

}

func revOrderInt(s []int) {
	fmt.Println("-----------------------revOrderInt-----------------------------------------")
	fmt.Printf("s before sorting:---> %v \n", s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Printf("s After sorting: --->%v \n", s)

}
