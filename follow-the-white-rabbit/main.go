package main

import (
	"errors"
	"fmt"
	"strings"
)

type Matrix2D [][]int

func printMatrix(matrix Matrix2D) {
	fmt.Println(strings.Repeat(" -", len(matrix[0])))
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func getDimensions(a Matrix2D) (int, int) {
	return len(a), len(a[0])
}

func haveEqualDimensions(a, b Matrix2D) bool {
	rowA, colA := getDimensions(a)
	rowB, colB := getDimensions(b)

	return rowA == rowB && colA == colB
}

func add(a, b Matrix2D) (Matrix2D, error) {

	if !haveEqualDimensions(a, b) {
		return nil, errors.New("the matrices have different dimensions")
	}

	rowCount, colCount := getDimensions(a)

	result := make(Matrix2D, rowCount)

	for i := 0; i < rowCount; i++ {
		result[i] = make([]int, colCount)
		for j := 0; j < colCount; j++ {
			result[i][j] = a[i][j] + b[i][j]
		}
	}

	return result, nil
}

func sub(a, b Matrix2D) (Matrix2D, error) {
	if !haveEqualDimensions(a, b) {
		return nil, errors.New("the matrices have different dimensions")
	}

	rowCount, colCount := getDimensions(a)

	result := make(Matrix2D, rowCount)

	for i := 0; i < rowCount; i++ {
		result[i] = make([]int, colCount)
		for j := 0; j < colCount; j++ {
			result[i][j] = a[i][j] - b[i][j]
		}
	}

	return result, nil
}

func multiply(a, b Matrix2D) (Matrix2D, error) {
	rowCountA, colCountA := getDimensions(a)
	rowCountB, colCountB := getDimensions(b)

	if colCountA != rowCountB {
		return nil, errors.New("the dimensions are not suitable for multiplying")
	}

	result := make(Matrix2D, rowCountA)

	for i := 0; i < rowCountA; i++ {
		result[i] = make([]int, colCountB)
		for j := 0; j < colCountB; j++ {
			for k := 0; k < colCountA; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result, nil

}

func main() {

	a := Matrix2D{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	printMatrix(a)

	b := Matrix2D{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{10, 11, 12, 13},
		{14, 15, 16, 17},
	}

	printMatrix(b)

	c, err := add(a, b)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("adding matrices")
		printMatrix(c)
	}

	d, err := sub(a, b)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("subtracting matrices")
		printMatrix(d)
	}

	fmt.Println("let's prove a + b = b + a")

	e, err := add(b, a)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("adding matrices")
		printMatrix(e)
	}

	fmt.Println("let's prove a - b != b - a")

	f, err := sub(b, a)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("subtracting matrices")
		printMatrix(f)
	}

	fmt.Println("let's multiply now!")

	g, err := multiply(a, b)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("multiplying matrices")
		printMatrix(g)
	}
}
