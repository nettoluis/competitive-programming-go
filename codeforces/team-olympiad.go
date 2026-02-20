package main

import "fmt"

func main() {
	var length int

	fmt.Scan(&length)
	gifted := make(map[int][]int)

	for i := 1; i <= length; i++ {
		var value int

		fmt.Scan(&value)

		gifted[value] = append(gifted[value], i)
	}
	
	counter := 0
	for j := 0; j < length; j++ {
		if (j < len(gifted[1]) && j < len(gifted[2]) && j < len(gifted[3])) {
			counter++
		}
	}
	fmt.Println(counter)

	saida := ""
	for k := 0; k < counter; k++ {
		if k != 0 {
			saida +=  "\n" + fmt.Sprintf("%d %d %d", gifted[1][k], gifted[2][k], gifted[3][k])
		} else {
			saida += fmt.Sprintf("%d %d %d", gifted[1][k], gifted[2][k], gifted[3][k])
		}
	}
	
	if counter != 0 {
		fmt.Println(saida)
	}

}
