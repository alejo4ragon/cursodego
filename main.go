package main

import (
	"fmt"
	"godesde0/variables"
)

func main() {
	// variables.ShowIntegers()
	// variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(342)
	fmt.Println(estado)
	fmt.Println(texto)
}
