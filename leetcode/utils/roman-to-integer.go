package utils

import "fmt"

const (
valores = map[string]int64 {
    "I" = 1,
    "V" = 5,
    "X" = 10,
    "L" = 50,
    "C" = 100,
    "D" = 500,
    "M" = 1000,
}
)

func romanToInt(s string) int {
    valores = map[string]int64{"I" = 1,"V" = 5,"X" = 10,"L" = 50,"C" = 100, "D" = 500,"M" = 1000}
    saida := 0
    for i := 0; i < len(s); i++ {
        atual, posterior := valores[string(s[i])], valores[string(s[i + 1])]
        if atual < posterior {
            saida -= atual
        } else {
            saida += atual
        }
    }

    saida += posterior

    return saida
}
