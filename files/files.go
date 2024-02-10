package files

import (
	"bufio"
	"fmt"
	"godesde0/ejercicios"
	"io/ioutil"
	"os"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.PedirNumeroGrabandoEnArchivo()

	archivo, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)

	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.PedirNumeroGrabandoEnArchivo()

	if !append(texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func append(texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Append: " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString: " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeerArchivo() {
	// Deprecado
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error durante el LeerArchivo: " + err.Error())
	}

	fmt.Println(string(archivo))
}

func LeerArchivoOs() {

	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error durante el LeerArchivoOs: " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
