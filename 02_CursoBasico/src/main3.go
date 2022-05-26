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
}
