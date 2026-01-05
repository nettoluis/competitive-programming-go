package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "unicode"
    "unicode/utf8"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    entrada, _ := reader.ReadString('\n')
    entrada = strings.TrimSpace(entrada)

    if entrada == "" {
        return
    }

    r, size := utf8.DecodeRuneInString(entrada)

    if unicode.IsUpper(r) {
        fmt.Println(entrada)
        return
    }

    nova := string(unicode.ToUpper(r)) + entrada[size:]

    fmt.Println(nova)
}
