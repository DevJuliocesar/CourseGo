package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaración de variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("suma:", result)

	// Resta
	result = x - y
	fmt.Println("resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Printf
	nombre := "Platzi"
	curso := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, curso)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, curso)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, curso)
	fmt.Println(message)

	// Tipo de dato
	fmt.Printf("nombre: %T\n", nombre)
	fmt.Printf("curso: %T", curso)
}
