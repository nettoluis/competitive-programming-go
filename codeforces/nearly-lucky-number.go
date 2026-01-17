package main

import "fmt"

func main() {
    var numero string

    fmt.Scan(&numero)

    contador := 0
    for _, v := range numero {
        if string(v) == "4" || string(v) == "7" {
            contador++
        }
    }

    if contador == 4 || contador == 7 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }

}
