package main

import (
    "fmt"
)

const EFICIENCIA = 12.0

func main() {
    var horas, velocidadeMedia float64
    
    fmt.Scan(&horas, &velocidadeMedia)

    fmt.Printf("%.3f\n", (horas * velocidadeMedia) / EFICIENCIA)
}
