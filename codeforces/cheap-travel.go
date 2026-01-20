package main

import "fmt"

func main() {
    var viagensDesejadas, viagensPacote, precoUnico, precoPacote int
    
    fmt.Scan(&viagensDesejadas, &viagensPacote, &precoUnico, &precoPacote)

    if viagensPacote * precoUnico <= precoPacote {
        fmt.Println(viagensDesejadas * precoUnico)
        return
    }

    qtdPacotes := viagensDesejadas / viagensPacote
    restante := viagensDesejadas - (viagensPacote * qtdPacotes)

    if precoUnico * restante < precoPacote {
        fmt.Println((precoUnico * restante) + (qtdPacotes * precoPacote))
    } else {
        fmt.Println((qtdPacotes + 1) * precoPacote)
    }
}
