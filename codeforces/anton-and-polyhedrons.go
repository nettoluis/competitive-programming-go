package main

import "fmt"


func main() {
	faces := map[string]int{"Tetrahedron" : 4, "Cube" : 6, "Octahedron" : 8, "Dodecahedron" : 12, "Icosahedron" : 20,}
	var n int

	fmt.Scan(&n)

	totalFaces := 0
	for i := 0; i < n; i++ {
		var polyhedron string

		fmt.Scan(&polyhedron)

		totalFaces += faces[polyhedron]
	}

	fmt.Println(totalFaces)
}
