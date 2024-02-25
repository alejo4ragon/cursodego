package funciones

import "fmt"

/*
* Las funciones anonimas fueron pensadas para
* no tener que crear muchas funciones, si no poder
* generar calculos internos dentro de una función
* junto con parámetros con los que ya cuenta esta
* función padre.
 */
func Calculos() {

	var numero3 int = 32
	var numero4 int = 243

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3 * numero4
	}

	fmt.Println(calculo(10, 25))

}
