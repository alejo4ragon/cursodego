package ejercicios

import (
	"strconv"
)

func Ejercicio1(texto string) (numero int, esNumero bool, resultado string) {
	numero, err := strconv.Atoi(texto)

	if err != nil {
		// Error al convertir el texto a número, no es un número válido
		return 0, false, "No es un número válido: " + err.Error()
	}

	if numero >= 100 {
		return numero, true, "Es mayor o igual a 100"
	}

	return numero, true, "Es menor a 100"
}
