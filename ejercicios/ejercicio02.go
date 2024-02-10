package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Numero int
var Err error

func PedirNumero() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingresa un número: ")

	for {
		if scanner.Scan() {
			ValorIngresado := scanner.Text()
			Numero, Err = strconv.Atoi(scanner.Text())
			if Err != nil {
				fmt.Println("Ingresa un número entero, el dato que has ingresado es incorrecto: " + ValorIngresado)
				continue
			} else {
				break
			}
		}
	}

	hasta := 10
	if Numero > 10 {
		hasta = Numero
	}

	for i := 1; i <= hasta; i++ {
		// multiplicacion := strconv.Itoa(Numero * i)
		// leyenda := strconv.Itoa(Numero) + "x" + strconv.Itoa(i) + " = " + multiplicacion
		// fmt.Println(leyenda)

		// Una sola línea
		fmt.Printf("%d x %d = %d \n", Numero, i, Numero*i)
	}

}
