package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Julio"] = 14
	m["Cesar"] = 20

	fmt.Println(m)

	value, ok := m["Julio"]
	fmt.Println(value, ok)
}
