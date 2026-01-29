package main

import "fmt"

func main() {
    var x int

    fmt.Scan(&x)

    formaBinaria := fmt.Sprintf("%b", x)

    digitos := make(map[rune]int)

    for i := range formaBinaria {
        digitos[rune(formaBinaria[i])]++
    }

    fmt.Println(digitos['1'])
}
