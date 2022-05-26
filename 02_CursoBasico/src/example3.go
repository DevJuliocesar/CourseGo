package main

import "fmt"

func main() {
	// For conditional
	for i := 0; i <= 2; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// For while
	counter := 0
	for counter < 2 {
		fmt.Println(counter)
		counter++
	}

	// Operador Logico
	myBool := true
	fmt.Println(!myBool)

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
}
