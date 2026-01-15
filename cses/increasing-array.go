package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) 

	
	scanInt := func() int {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		return val
	}

	
	if !scanner.Scan() {
		return
	}
	qtd, _ := strconv.Atoi(scanner.Text())

	
	numeros := make([]int, qtd)
	for i := 0; i < qtd; i++ {
		numeros[i] = scanInt()
	}

	
	var moves int64 = 0
	maiorAtual := 0

	for _, v := range numeros {
		if maiorAtual < v {
			maiorAtual = v
		} else {
			moves += int64(maiorAtual - v)
		}
	}

	fmt.Println(moves)
}
