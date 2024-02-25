package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "Ciudad de méxico"
	paises["Argentina"] = "Buenos Aires"
	paises["Colombia"] = "Bogotá"
	paises["España"] = "Madrid"
	paises["Alemania"] = "Berlín"

	// fmt.Println(paises)
	// fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	// fmt.Println(campeonato)

	// for equipo, puntaje := range campeonato {
	// 	fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	// }

	delete(campeonato, "Chivas")
	fmt.Println(campeonato)

	campeonato["Juventus"] = 80

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe = %t \n", puntaje, existe)
	fmt.Println(campeonato)

}
