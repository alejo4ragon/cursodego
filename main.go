package main

import "godesde0/middleware"

func main() {
	// variables.ShowIntegers()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoaTexto(342)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// Condicionales
	// os := runtime.GOOS

	// fmt.Println(os)

	// if os := runtime.GOOS; os == "linux" || os == "OS X." || os == "darwin" {
	// 	fmt.Println("Esto no es Windows, es ", os)
	// } else {
	// 	fmt.Println("Esto es Windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// numero, esNumero, resultado := ejercicios.Ejercicio1("1asdasd")
	// fmt.Println(numero, esNumero, resultado)

	// Recibir datos por teclado
	// teclado.IngresoNumeros()

	// Iteraciones
	// iteraciones.IterarPasosGrandes(0, 5, 100)

	// ejercicios.PedirNumero()
	// fmt.Println(ejercicios.PedirNumeroGrabandoEnArchivo())

	// files.GrabaTabla()
	// files.SumaTabla()
	// files.LeerArchivo()
	// files.LeerArchivoOs()

	// funciones.LlamarClosure()
	// funciones.Exponencia(10)
	// arreglos_slices.MuestroArreglos()
	// arreglos_slices.MuestroMatriz()
	// arreglos_slices.MuestroSlice()
	// arreglos_slices.Capacidad()
	// mapas.MostrarMapas()
	// users.AltaUsuario()

	// Johan := new(modelos.Hombre)
	// ejer_interfaces.HumanosRespirando(Johan)

	// Geral := new(modelos.Mujer)
	// ejer_interfaces.HumanosRespirando(Geral)
	// defer_panic.VemosDefer()
	// defer_panic.EjemploPanic()

	// canal1 := make(chan bool)
	// go goroutines.MiNombreLentooo("Johan Alejandro Aragón Peña", canal1)

	// fmt.Println("Estoy aquí")
	// defer func() {
	// 	<-canal1
	// }()

	// var x string
	// fmt.Scanln(&x)

	// webserver.MiWebServer()
	middleware.MiMiddleware()
}
