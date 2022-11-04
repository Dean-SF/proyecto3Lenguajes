package main

import (
	"fmt"
	gen "lenguajes/GeneradorPseudo-Aleatorios/modulos"
)

func main() {
	gen := gen.GeneradorParametros{Semilla: 11, N: 500, X: 2040, Limit: 500}

	// Cada vez que se hace Next se consigue el siguiente numero generado
	for {
		result := gen.Generador()
		if result == 0 {
			return
		}
		fmt.Printf("%v ", result)
	}
}
