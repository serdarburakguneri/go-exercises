package main

import "fmt"

func selectionSort(collection []int) {
	fmt.Println("sorting started")
	length := len(collection)
	for i := 0; i < length-1; i++ {
		index := i
		for k := i + 1; k < length; k++ {
			if collection[k] < collection[index] {
				index = k
			}
		}
		tmp := collection[i]
		collection[i] = collection[index]
		collection[index] = tmp
		fmt.Println(collection)
	}
	fmt.Println("sorting ended")
}

func main() {
	collection := []int{1, 5, 6, 3, 99, 44, 55, 22}
	fmt.Println("before sorting")
	fmt.Println(collection)
	selectionSort(collection)
	fmt.Println("after sorting")
	fmt.Println(collection)
}
