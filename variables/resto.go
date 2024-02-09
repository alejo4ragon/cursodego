package variables

import (
	"fmt"
	"strconv"
	"time"
)

// Variables globales se pueden usar en cualquier paquete
// si las variables están con la primera letra en mayúscula
var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Johan"
	Estado = true
	Sueldo = 3500.66
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
