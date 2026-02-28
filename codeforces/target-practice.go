package main

import "fmt"

const (
	MATRIX_ORDER = 10
	HIGHEST_SCORE = 5
)

func runBorders(k int, m []string) int {
	counter := 0

	for j := k; j < MATRIX_ORDER-1-k; j++ {
		if string(m[k][j]) == "X" {
			counter += k+1
		}
	}

	for i := k; i < MATRIX_ORDER-1-k; i++ {
		if string(m[i][MATRIX_ORDER-1-k]) == "X" {
			counter += k+1
		}
	}

	for j := MATRIX_ORDER-1-k; j > k; j-- {
		if string(m[MATRIX_ORDER-1-k][j]) == "X" {
			counter += k+1
		}
	}


	for i := MATRIX_ORDER-1-k; i > k; i-- {
		if string(m[i][k]) == "X" {
			counter += k+1
		}
	}

	return counter
}

func main() {
	var testCases int

	fmt.Scan(&testCases)

	for t := 0; t < testCases; t++ {
		matrix := make([]string, MATRIX_ORDER)

		for i := 0; i < len(matrix); i++ {
			fmt.Scan(&matrix[i])
		}

		score := 0
		for k := 0; k < HIGHEST_SCORE; k++ {
			score += runBorders(k, matrix)
		}

		fmt.Println(score)
	}
}
