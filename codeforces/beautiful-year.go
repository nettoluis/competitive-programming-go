package main

import "fmt"

func ehBonito(ano string) bool {
    unicos := make(map[string]bool)

    for _, valor := range ano {
        unicos[string(valor)] = true
    }

    return len(unicos) == 4
}

func main() {
    var ano int

    fmt.Scan(&ano)
    ano++

    for true {
        if ehBonito(fmt.Sprintf("%d", ano)) {
            break
        }
        ano++
    }

    fmt.Println(ano)
}
