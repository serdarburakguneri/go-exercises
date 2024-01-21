package main

import "fmt"

func bubleSort(collection []int) {
	length := len(collection)
	for i := 0; i < length-1; i++ {
		for k := 0; k < length-i-1; k++ {
			if collection[k] > collection[k+1] {
				//let's swap
				tmp := collection[k]
				collection[k] = collection[k+1]
				collection[k+1] = tmp
			}
		}
		fmt.Println(collection)
	}
}

func main() {
	collection := []int{1, 5, 6, 3, 99, 44, 55, 22}
	fmt.Println("before sorting")
	fmt.Println(collection)
	bubleSort(collection)
	fmt.Println("after sorting")
	fmt.Println(collection)
}
