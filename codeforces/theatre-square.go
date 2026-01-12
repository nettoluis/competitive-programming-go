package main

import (
    "fmt"
    "math"
)

func main() {
    var n, m, a float64

    fmt.Scan(&n, &m, &a)
    
    resultado := math.Ceil(n / a) * math.Ceil(m / a)

    fmt.Println(int(resultado))
}
