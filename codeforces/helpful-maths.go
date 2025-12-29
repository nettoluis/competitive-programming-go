package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
    "sort"
)

func parseIntegers(s []string) ([]int, error) {
	saida := make([]int, len(s))

	for i, strVal := range s {
		val, err := strconv.Atoi(strings.TrimSpace(strVal))
		if err != nil {
			return nil, fmt.Errorf("falha ao converter '%s': %w", strVal, err)
		}
		saida[i] = val
	}

	return saida, nil
}

func main() {
	var entrada string

	if _, err := fmt.Scan(&entrada); err != nil {
		log.Fatal(err)
	}

	nums := strings.Split(entrada, "+")

	numsInt, err := parseIntegers(nums)
	if err != nil {
		log.Fatalf("Erro na entrada: %v", err)
	}

    sort.Ints(numsInt)

	var sb strings.Builder
	for i, num := range numsInt {
		if i > 0 {
			sb.WriteString("+")
		}
		sb.WriteString(strconv.Itoa(num))
	}

	fmt.Println(sb.String())
}
