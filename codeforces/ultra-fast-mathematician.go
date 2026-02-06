package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1, s2 string

	fmt.Scan(&s1, &s2)

	saida := ""
	for i := 0; i < len(s1); i++ {
		digito1, err1 := strconv.Atoi(string(s1[i]))
		digito2, err2 := strconv.Atoi(string(s2[i]))

		if err1 != nil || err2 != nil {
			panic(err1)
		}

		saida += strconv.Itoa(digito1 ^ digito2)
	}

	fmt.Println(saida)
}
