package arreglos_slices

import "fmt"

var tabla [10]int //Arreglo de 10 elementos
func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)
}
