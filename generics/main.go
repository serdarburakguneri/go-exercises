package main

import "fmt"

func Filter[T any](collection []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func has6letters(name string) bool {
	return len(name) == 6
}

func even(number int) bool {
	return number%2 == 0
}

func main() {

	names := []string{"serdar", "burak", "guneri"}
	fmt.Println(names)
	filteredNames := Filter(names, has6letters) //or Filter[string](names, has6letters)
	fmt.Println(filteredNames)

	numbers := []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	fmt.Println(numbers)
	filteredNumbers := Filter(numbers, even)
	fmt.Println(filteredNumbers)

}
