package main

import (
    "fmt"
)

const MINUTOS_POR_QUILOMETRO = 2

func main() {
    var distancia int

    fmt.Scan(&distancia)

    fmt.Printf("%d minutos\n", distancia * MINUTOS_POR_QUILOMETRO)
}
