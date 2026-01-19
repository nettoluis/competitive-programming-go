package main

import (
    "fmt"
    "unicode"
)

func ehVogal(s string) bool {
    vogais := "aeiouyAEIOUY"

    for _, v := range vogais {
        if s == string(v) {
            return true
        }
    }

    return false
}

func main() {
    var palavra string

    fmt.Scan(&palavra)

    saida := ""
    for _, v := range palavra {
        if !ehVogal(string(v)) {
            saida += "." + string(unicode.ToLower(v))
        }
    }

    fmt.Println(saida)
}
