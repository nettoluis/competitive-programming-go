package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const LIMIT = 1000000 

var isPrime [LIMIT + 1]bool

func sieve() {
	
	for i := 2; i <= LIMIT; i++ {
		isPrime[i] = true
	}
	
	isPrime[0] = false
	isPrime[1] = false

	
	for p := 2; p*p <= LIMIT; p++ {
		if isPrime[p] {
			for i := p * p; i <= LIMIT; i += p {
				isPrime[i] = false
			}
		}
	}
}

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	sieve()

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var val int64
		fmt.Fscan(reader, &val)
		
		root := int64(math.Sqrt(float64(val)))

		if root*root == val && isPrime[root] {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}
