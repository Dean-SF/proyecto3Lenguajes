package main

import (
	"fmt"
	gen "lenguajes/GeneradorPseudo-Aleatorios/modulos"
)

func main() {
	gen := gen.Generador{Semilla: 11, N: 500, X: 2040, Limit: 500}

	// n se utiliza para ir restando de 1 en 1 a la cantidad de numeros a generar y parar el trabajo cuando se llega a 0
	// x será una constante

	/*
		limite es una constante con la cantidad de numeros a generar, se utiliza para identificar cuándo se está por generar
		el primer número, para así inicializar x con la semilla.
		Se debe ingresar el mismo número para n y límite
	*/

	// Cada vez que se hace Next se consigue el siguiente numero generado
	for {
		result := gen.Next()
		if result == nil {
			return
		}
		fmt.Printf("%v ", *result)
	}
}
