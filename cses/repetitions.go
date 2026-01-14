package main

import "fmt"

func main() {
    var sequencia string

    fmt.Scan(&sequencia)

    contador := 1
    contadorMax := 0
    for i := 1; i < len(sequencia); i++ {
        if sequencia[i] == sequencia[i - 1] {
            contador++
        } else {
            if contador > contadorMax {
                contadorMax = *&contador
            }
            contador = 1
        }
    }

    if contador > contadorMax {
        contadorMax = *&contador
    }

    fmt.Println(contadorMax)
}
