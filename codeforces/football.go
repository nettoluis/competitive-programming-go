package main

import "fmt"

func janelaDeslizante(inicio int, elemento byte, sequencia string) bool {
    for i := inicio; i < inicio+7; i++ {
        if sequencia[i] != elemento {
            return false
        }
    }
    return true
}

func main() {
    var jogo string
    fmt.Scan(&jogo)

    encontrado := false
    for i := 0; i+7 <= len(jogo); i++ {
        if jogo[i] == '1' {
            encontrado = janelaDeslizante(i, '1', jogo)
        } else {
            encontrado = janelaDeslizante(i, '0', jogo)
        }
        
        if encontrado {
            break
        }
    }

    if encontrado {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
