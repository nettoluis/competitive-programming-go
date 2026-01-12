package main

import (
    "fmt"
    "strings"
    "unicode"
)

func main() {
    var palavra string

    fmt.Scan(&palavra)

    maiusculas, minusculas := 0, 0
    for _, letra := range palavra {
        if unicode.IsLower(letra) {
            minusculas++
        } else {
            maiusculas++
        }
    }

    if minusculas >= maiusculas {
        fmt.Println(strings.ToLower(palavra))
    } else {
        fmt.Println(strings.ToUpper(palavra))
    }
}
