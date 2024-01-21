package main

import "fmt"

func insertionSort(collection []int) {
	fmt.Println("sorting started")
	length := len(collection)
	for i := 1; i < length; i++ {
		current := collection[i]
		j := i - 1
		for j >= 0 && collection[j] > current {
			collection[j+1] = collection[j]
			j--
		}
		collection[j+1] = current
		fmt.Println(collection)
	}
	fmt.Println("sorting ended")
}

func main() {
	collection := []int{1, 5, 6, 3, 99, 44, 55, 22}
	fmt.Println("before sorting")
	fmt.Println(collection)
	insertionSort(collection)
	fmt.Println("after sorting")
	fmt.Println(collection)
}
