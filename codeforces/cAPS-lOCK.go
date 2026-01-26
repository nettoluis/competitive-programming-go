package main

import (
    "fmt"
    "unicode"
)

func todaEmCaps(palavra string) bool {
    for _, v := range palavra {
        if unicode.IsLower(v) {
            return false
        }
    }

    return true
}

func apenasPrimeiraLetraMinuscula(palavra string) bool {
    if unicode.IsUpper(rune(palavra[0])) {
        return false
    }

    for i := 1; i < len(palavra); i++ {
        if unicode.IsLower(rune(palavra[i])) {
            return false
        }
    }

    return true
}

func corrigindoPalavra(palavra []rune) {
    for i, _ := range palavra {
        if palavra[i] >= 97 {
            palavra[i] -= 32
        } else {
            palavra[i] += 32
        }
    }
}

func main() {
    var palavra string
    
    fmt.Scan(&palavra)

    palavraCorrigida := []rune(palavra)
    if todaEmCaps(palavra) || apenasPrimeiraLetraMinuscula(palavra) {
        corrigindoPalavra(palavraCorrigida)
    }

    fmt.Println(string(palavraCorrigida))
}
