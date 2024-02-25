package arreglos_slices

import "fmt"

var tabla [10]int //Arreglo de 10 elementos
var tabla2 [10]int = [10]int{
	1,
	2,
	3,
	4,
}

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	// fmt.Println(tabla)
	// fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	arreglo := [10]string{"Johan", "Alejandro", "Aragón"}

	for i := 0; i < len(arreglo); i++ {
		fmt.Println(arreglo[i])
	}
}

var matriz [20][30]int

func MuestroMatriz() {
	largo := len(matriz)
	ancho := len(matriz[0])

	contador := 0

	for i := 0; i < largo; i++ {
		for j := 0; j < ancho; j++ {
			matriz[i][j] = contador
			contador++
		}
	}

	fmt.Println(matriz)
}
