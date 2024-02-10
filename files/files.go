package files

import (
	"fmt"
	"godesde0/ejercicios"
	"os"
)

func GrabaTabla() {
	var texto string = ejercicios.PedirNumeroGrabandoEnArchivo()

	archivo, err := os.Create("./files/txt/tabla.txt")

	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)

	archivo.Close()
}

// func SumaTabla() {
// 	var texto string = ejercicios.PedirNumeroGrabandoEnArchivo()
// }
