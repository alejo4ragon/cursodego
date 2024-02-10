package iteraciones

import (
	"fmt"
)

/**
* Solo podemos usar for para iterar
* para simular un while se usa
* for {
*
*}
* usar break para romper el ciclo, se usa continue para evitar algo
**/

func Iterar() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func IterarMenosLineas() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func IterarPasosGrandes(desde int, paso int, hasta int) {

	for i := desde; i <= hasta; i += paso {
		fmt.Println(i)
	}
}
