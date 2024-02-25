package arreglos_slices

import "fmt"

var tabla [10]int //Arreglo de 10 elementos
var tabla2 [10]int = [10]int{1, 2, 3, 4}

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)
	fmt.Println(tabla2)
}
