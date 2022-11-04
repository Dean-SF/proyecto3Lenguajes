package main

import (
	"fmt"
	gen "lenguajes/GeneradorPseudo-Aleatorios/modulos"
)

func main() {
	cant := 1500
	semilla := 47

	gen := gen.GeneradorParametros{Semilla: semilla, N: cant, X: 2040, Limit: cant}

	// Cada vez que se llama al método se consigue el siguiente numero generado
	for {
		result := gen.Generador()
		if gen.N == 0 {
			return
		}
		fmt.Printf("%v ,", result)
	}
}
