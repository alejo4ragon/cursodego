package main

import (
	"godesde0/teclado"
)

func main() {
	// variables.ShowIntegers()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoaTexto(342)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// Condicionales
	// os := runtime.GOOS

	// fmt.Println(os)

	// if os := runtime.GOOS; os == "linux" || os == "OS X." || os == "darwin" {
	// 	fmt.Println("Esto no es Windows, es ", os)
	// } else {
	// 	fmt.Println("Esto es Windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// numero, esNumero, resultado := ejercicios.Ejercicio1("1asdasd")
	// fmt.Println(numero, esNumero, resultado)

	// Recibir datos por teclado
	teclado.IngresoNumeros()
}
