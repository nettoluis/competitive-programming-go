package main

import "fmt"

func abreviaPalavra(palavra string) string {
    var saida string
    var primeiraLetra string = string(palavra[0])
    var ultimaLetra string = string(palavra[len(palavra) - 1])
    var qtdLetrasIntermediarias int = len(palavra) - 2

    saida = fmt.Sprintf("%s%d%s", primeiraLetra, qtdLetrasIntermediarias, ultimaLetra)

    return saida
}


func main() {
    var qtd int
    fmt.Scanf("%d\n", &qtd)

    for i := 0; i < qtd; i++ {
       var palavra string
       fmt.Scanln(&palavra)

       if tamanho := len(palavra); tamanho > 10 {
            fmt.Println(abreviaPalavra(palavra))
       } else {
            fmt.Println(palavra)
       }
    }
}
