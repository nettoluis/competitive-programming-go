package main

import (
    "fmt"
    "strings"
)

func saoIguais(primeira, segunda string) bool {
    for i := 0; i < len(primeira); i++ {
        if primeira[i] != segunda[i] {
            return false
        }
    }

    return true
}


func main() {
    var primeira, segunda string
    fmt.Scan(&primeira, &segunda)

    primeira = strings.ToLower(primeira)
    segunda = strings.ToLower(segunda)

    switch {
        case saoIguais(primeira, segunda):
            fmt.Println("0")
        case primeira < segunda:
            fmt.Println("-1")
        case primeira > segunda:
            fmt.Println("1")
    }
}
