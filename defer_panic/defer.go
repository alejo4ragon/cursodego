package defer_panic

import (
	"fmt"
	"log"
)

/*
* En primera instancia defer sirve para ejecutar al final un comando
 */
func VemosDefer() {
	fmt.Println("Este es un primer mensaje")

	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {

	defer func() {
		reco := recover()
		if reco != nil { // si reco es distinto de nil significa que hubo un panic
			log.Fatalf("ocurrió un error que generó Panic: \n %v", reco)
		}
	}()

	a := 1

	if a == 1 {
		panic("A no puede ser igual a 1")
	}
}
