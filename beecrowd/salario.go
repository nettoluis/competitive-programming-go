package main

import "fmt"

func main() {
    var numero int
    var horas, salarioPorHora float64

    fmt.Scan(&numero, &horas, &salarioPorHora)

    fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", numero, horas * salarioPorHora)
}
