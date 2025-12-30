package main

import "fmt"

func main() {
    var nota1, nota2, nota3, nota4 float64
    fmt.Scan(&nota1, &nota2, &nota3, &nota4)

    media1 := ((nota1 * 2) + (nota2 * 3) + (nota3 * 4) + (nota4 * 1)) / 10

    fmt.Printf("Media: %.1f\n", media1)
    switch {
        case media1 >= 7.0:
            fmt.Println("Aluno aprovado.")
        case media1 < 5.0:
            fmt.Println("Aluno reprovado.")
        default:
            fmt.Println("Aluno em exame.")

            var nota5 float64
            fmt.Scan(&nota5)

            fmt.Printf("Nota do exame: %.1f\n", nota5)
            
            media2 := (media1 + nota5) / 2
            
            switch {
                case media2 >= 5.0:
                    fmt.Println("Aluno aprovado.")
                case media2 < 5.0:
                    fmt.Println("Aluno reprovado.")
            }

            fmt.Printf("Media final: %.1f\n", media2)
    }
}
