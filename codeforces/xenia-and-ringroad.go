package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    var tamanhoDaRua, qtdTarefas int
    
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    fmt.Fscan(reader, &tamanhoDaRua, &qtdTarefas)


    var tempo int64 = 0
    var posicaoAtual int = 1
    for j := 0; j < qtdTarefas; {
        var objetivo int

        fmt.Fscan(reader, &objetivo)

        if objetivo >= posicaoAtual {
            tempo += int64(objetivo - posicaoAtual)
        } else {
            tempo += int64((tamanhoDaRua - posicaoAtual) + objetivo)
        }

        posicaoAtual = objetivo
        j++
    }

    fmt.Fprintln(writer, tempo)
}
