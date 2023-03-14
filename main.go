package main

import "fmt"

func main() {

	//ex 1

	var nome string
	fmt.Print("Qual é o seu nome? ")
	fmt.Scan(&nome)
	fmt.Println("Olá,", nome)

	var idade int64
	fmt.Print("Quantos anos você tem?")
	fmt.Scan(&idade)
	fmt.Println("Você tem", idade)

	var peso float64
	fmt.Print("Quanto você pesa?")
	fmt.Scan(&peso)
	fmt.Println("Você pesa", peso)

	//ex 2

	var base float64
	fmt.Print("Informe a base do retângulo.")
	fmt.Scan(&base)
	fmt.Println("A base do retângulo é:", base)

	var altura float64
	fmt.Print("Informe a altura do retângulo.")
	fmt.Scan(&altura)
	fmt.Println("A altura do retângulo é:", altura)

	var area float64 = base * altura
	fmt.Println("A área do retângulo é:", area)
}
