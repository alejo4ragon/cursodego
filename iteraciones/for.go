package iteraciones

import (
	"fmt"
)

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
