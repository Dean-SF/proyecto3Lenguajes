package gen

import (
	"fmt"
	"os"
)

// Funcion auxiliar para determinar si el numero recibido es primo
func EsPrimo(num int) bool {
	// Calcular el residuo de la división entre el número dado por parámetro y los números del 2 hasta el mismo
	for i := 2; i <= num; i++ {
		if num%i == 0 {
			return false //  No es primo
		} else {
			return true // Sí es primo
		}
	}
	return true
}

/*
N se utiliza para ir restando de 1 en 1 a la cantidad de numeros a generar
y parar el trabajo cuando se llega a 0.

X será una constante del parámetro x para la fórmula, después del primer número
generado va a almacenar el número generado anteriormente.

Limite es una constante con la cantidad de numeros a generar, se utiliza para identificar cuándo
se está por generar el primer número, para así inicializar x con la semilla.
Se debe ingresar el mismo número para n y límite.
*/
type GeneradorParametros struct {
	Semilla, N, X, Limit int
}

func (gen *GeneradorParametros) Generador() int {
	// --------- Validaciones ----------------
	if gen.Limit < 500 || gen.Limit > 5000 {
		error := "El limite y n deben ser un numero entero entre 500 y 5000"
		fmt.Println(error)
		os.Exit(1)
	}

	if gen.Semilla < 11 || gen.Semilla > 257 {
		error := "La semilla debe ser un numero entero primo entre 11 y 257"
		fmt.Println(error)
		os.Exit(2)
	}

	if EsPrimo(gen.Semilla) == false {
		error := "La semilla debe ser un numero entero primo entre 11 y 257"
		fmt.Println(error)
		os.Exit(3)
	}

	// Si gen.N es igual a 0, significa que hemos generado todos los números
	if gen.N == 0 {
		return 0
	}

	/*
		Si gen.N es igual a gen.Limit, significa que vamos a generar el primer número,
		entonces inicializamos x  con el valor de la semilla
	*/
	if gen.N == gen.Limit {
		gen.X = gen.Semilla
	}

	m := 2048 // período
	a := 109  // multiplicador
	b := 853  // incremento

	result := (a*gen.X + b) % m
	gen.X = result        // actualizar el valor de X
	result = result % 255 // Convertir al intervalo [0,255]

	gen.N-- // Restarle 1 a N para indicar que hemos generado 1 número
	return result
}
