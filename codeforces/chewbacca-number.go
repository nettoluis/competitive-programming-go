package main

import (
    "fmt"
    "strconv"
)

func main() {
    var numero string

    fmt.Scan(&numero)

    saida := ""
    for i := 0; i < len(numero); i++ {
        numeroInt, err := strconv.Atoi(string(numero[i]))
        if err != nil {
            panic(err)
        }

        if numeroInt < 5 || (numeroInt == 9 && i == 0) {
            saida += fmt.Sprintf("%d", numeroInt)
        } else {
            saida += fmt.Sprintf("%d", 9 - numeroInt)
        }
    }

    fmt.Println(saida)
}
