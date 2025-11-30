package main

import "fmt"

func main() {
    var nome string
    var salarioFixo, vendas float64

    fmt.Scan(&nome, &salarioFixo, &vendas)

    fmt.Printf("TOTAL = R$ %.2f\n", salarioFixo + (0.15 * vendas))
}
