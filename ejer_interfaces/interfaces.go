package ejer_interfaces

import (
	"fmt"
	"godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}

func SerVivoComiendo(sv interfaces.Humano) {
	sv.Comer()
	sv.EstaVivo()
	fmt.Printf("Soy un/a %s, y estoy comiendo %t \n", sv.Sexo())
}
