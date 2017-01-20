package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numPesi   = 3
	numMonete = 9999999
	casuale   = 100
)

var indiceMoney int

// main is the entry point for the application.
func main() {
	var copyMonete [numMonete]int
	//monete[3] = casuale

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randIndice := r1.Intn(numMonete)
	//isDestra := false
	monete := make([]int, numMonete)
	monete[randIndice] = casuale

	fmt.Printf("%v   indice = %d \n", monete, randIndice)

	copy(copyMonete[:], monete)
	indiceMoney = probMod(monete)
	//	posMoneta := findMoney(monete, 3)
	fmt.Printf("%v : la moneta nella posizione [%d] di valore [ %d ] e quella falsa \n ", copyMonete, indiceMoney, copyMonete[indiceMoney])
}

func findMoney(money []int, pos int) int {

	dimSlice := len(money)
	var sliceD []int
	var posizione int
	//fmt.Println(dimSlice)

	totFirstPart, totSecontPart := sumSlice(money, dimSlice/2)
	//fmt.Println("totale monete Prima Parte: ", totFirstPart, "totale monete Seconda Parte: ", totSecontPart)
	if totFirstPart > totSecontPart {
		sliceD = money[:3]
		fmt.Println("Prima parte: ", sliceD)
		posizione = swap(money)
	} else {
		sliceD = money[3:]
		posizione = swapB(money)
		fmt.Println("seconda parte: ", sliceD)
	}
	//posizione = swapD(sliceD)
	return posizione
}

func sumSlice(slice []int, dimSlice int) (totFirstPart, totSecontPart int) {

	for i, valore := range slice {
		if i < dimSlice {
			totFirstPart += valore

		} else {

			totSecontPart += valore
		}
	}

	return totFirstPart, totSecontPart
}

func swap(slice []int) int {

	var posMoneta int

	if slice[0]+slice[3] == slice[1]+slice[4] {
		if slice[2] != 0 {
			posMoneta = 2
		} else {
			posMoneta = 5
		}

	} else if slice[0]+slice[3] > slice[1]+slice[4] {
		if slice[0] != 0 {
			posMoneta = 0
		} else {
			posMoneta = 4
		}
	} else {
		if slice[1] != 0 {
			posMoneta = 1
		} else {
			posMoneta = 3
		}
	}
	return posMoneta
}

func swapB(slice []int) int {

	var posMoneta int

	if slice[0]+slice[3] == slice[1]+slice[4] {
		if slice[2] != 0 {
			posMoneta = 2
		} else {
			posMoneta = 5
		}

	} else if slice[0]+slice[3] > slice[1]+slice[4] {
		if slice[1] != 0 {
			posMoneta = 1
		} else {
			posMoneta = 3
		}
	} else {
		if slice[0] != 0 {
			posMoneta = 0
		} else {
			posMoneta = 4
		}
	}
	return posMoneta
}

func swapD(slice []int) int {

	var posMoneta int

	if slice[0]+slice[3] == slice[1]+slice[4] {
		if slice[2] != 0 {
			posMoneta = 2
		} else {
			posMoneta = 5
		}

	} else if slice[0]+slice[3] > slice[1]+slice[4] {
		if slice[0] != 0 {
			posMoneta = 0
		} else {
			posMoneta = 4
		}
	} else {
		if slice[1] != 0 {
			posMoneta = 1
		} else {
			posMoneta = 3
		}
	}
	return posMoneta
}

func probMod(money []int) int {

	if len(money) == 1 {
		return 0
	}

	sliceD, isDestra := bilancia(money)

	indiceMoney := probMod(sliceD)
	//fmt.Println("...............................................")
	//	fmt.Printf("sliceD= %v\n", sliceD)
	//fmt.Printf("money= %v\n", money)

	if isDestra {
		i := len(money) / 2
		indiceMoney += i
	}

	return indiceMoney
}

func bilancia(money []int) ([]int, bool) {
	var destra bool
	if pesaMoney(money[:len(money)/2]) > pesaMoney(money[len(money)/2:]) {

		return money[:len(money)/2], destra
	}

	destra = true
	return money[len(money)/2:], destra

}

func pesaMoney(money []int) int {

	var tot int
	for _, val := range money {
		tot += val
	}
	return tot
}
