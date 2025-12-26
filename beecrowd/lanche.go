package main

import "fmt"

func main() {
    precos := map[int]float64{
        1:4.00,
        2:4.50,
        3:5.00,
        4:2.00,
        5:1.50,
    }

    var codigo, quantidade int 
    fmt.Scan(&codigo, &quantidade)

    fmt.Printf("Total: R$ %.2f\n", precos[codigo] * float64(quantidade))
}
