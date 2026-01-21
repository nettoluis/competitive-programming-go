package main

import (
    "fmt"
)

const (
    meta = "hello"
)

func consegueFormar(palavra, meta string) string {
    contador := 0
    for i := 0; i < len(palavra); i++ {
        if contador < 5 && palavra[i] == meta[contador] {
            contador++
        }
    }    

    if contador == 5 {
        return "YES"
    } else {
        return "NO"
    }
}

func main() {
    var palavra string

    fmt.Scan(&palavra)

    saida := consegueFormar(palavra, meta)

    fmt.Println(saida)
}
