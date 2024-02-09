package variables

import (
	"fmt"
)

func ShowIntegers() {
	var intCommon int
	intOf32 := int32(10)
	intOf64 := int64(100)

	fmt.Println("Entero común: ", intCommon)
	fmt.Println("Entero de 32: ", intOf32)
	fmt.Println("Entero de 64: ", intOf64)
}
