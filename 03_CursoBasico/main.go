package main

import (
	pk "cursoBasico/mypackage"
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hola")

	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(*b)

	myPc := pc{ram: 16, disk: 1000, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)
}
