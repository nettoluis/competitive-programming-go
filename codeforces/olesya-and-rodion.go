package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var length, divisor int

	fmt.Scan(&length, &divisor)

	if divisor != 10 {
		fmt.Println(strconv.Itoa(divisor) + strings.Repeat("0", length - 1))
	} else if length != 1 {
		fmt.Println(strings.Repeat("1", length - 1) + "0")
	} else {
		fmt.Println(-1)
	}
}
