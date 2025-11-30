package main

import "fmt"

func main() {
    var kilos int;
    fmt.Scanf("%d\n", &kilos)

    if (kilos > 2) && (kilos % 2 == 0) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
