package main

import "fmt"

func main() {

	//ex 1

	var nome string
	fmt.Print("Qual é o seu nome? ")
	fmt.Scan(&nome)
	fmt.Println(nome, "é um nome muito bonito!")

	var idade int64
	fmt.Print("Quantos anos você tem? ")
	fmt.Scan(&idade)
	fmt.Println("Você tem", idade, "anos.")

	var peso float64
	fmt.Print("Quanto você pesa? ")
	fmt.Scan(&peso)
	fmt.Println("Você pesa", peso, "kgs.")

	//ex 2

	var base float64
	fmt.Print("Informe a base do retângulo na unidade de medida dos centímetros: ")
	fmt.Scan(&base)
	fmt.Println("A base do retângulo é:", base, "cms.")

	var altura float64
	fmt.Print("Informe a altura do retângulo na unidade de medida dos centímetros: ")
	fmt.Scan(&altura)
	fmt.Println("A altura do retângulo é:", altura, "cms.")

	var área float64 = base * altura
	fmt.Println("A área do retângulo é:", área, "cms.")

	//ex 3

	var base_caixa float64
	fmt.Print("Informe a base da caixa na unidade de medida dos centímetros: ")
	fmt.Scan(&base_caixa)
	fmt.Println("A base da caixa é:", base_caixa, "cms.")

	var altura_caixa float64
	fmt.Print("Informe a altura da caixa na unidade de medida dos centímetros: ")
	fmt.Scan(&altura_caixa)
	fmt.Println("A altura da caixa é:", altura_caixa, "cms.")

	var profundidade_caixa float64
	fmt.Print("Informe a profundidade da caixa na unidade de medida dos centímetros: ")
	fmt.Scan(&profundidade_caixa)
	fmt.Println("A profundidade da caixa é:", profundidade_caixa, "cms.")

	var volume_caixa float64 = base_caixa * altura_caixa * profundidade_caixa
	fmt.Println("O volume da caixa é:", volume_caixa, "cms.")

	//ex 4

	var primeiro_salário float64
	fmt.Print("Informe o primeiro salário: ")
	fmt.Scan(&primeiro_salário)
	fmt.Println("O primeiro salário foi de", primeiro_salário, "reais.")

	var segundo_salário float64
	fmt.Print("Informe o segundo salário: ")
	fmt.Scan(&segundo_salário)
	fmt.Println("O segundo salário foi de", segundo_salário, "reais.")

	var terceiro_salário float64
	fmt.Print("Informe o terceiro salário: ")
	fmt.Scan(&terceiro_salário)
	fmt.Println("O terceiro salário foi de", terceiro_salário, "reais.")

	var quarto_salário float64
	fmt.Print("Informe o quarto salário: ")
	fmt.Scan(&quarto_salário)
	fmt.Println("O quarto salário foi de", quarto_salário, "reais.")

	var média_salários float64 = (primeiro_salário + segundo_salário + terceiro_salário + quarto_salário) / 4
	fmt.Println("A média salarial nesse período de tempo foi de", média_salários, "reais.")

	//ex 5

	var dinheiro_reais float64
	fmt.Print("Informe a quantidade de dinheiro em reais: ")
	fmt.Scan(&dinheiro_reais)
	fmt.Println("A quantidade de dinheiro que você tem é de", dinheiro_reais, "reais.")

	var cotação_dólar_atual float64
	fmt.Print("Informe com quantos reais pode-se comprar com 1 dólar: ")
	fmt.Scan(&cotação_dólar_atual)
	fmt.Println("Com 1 dolár, pode-se comprar", cotação_dólar_atual, "reais.")

	var dinheiro_dólares float64 = dinheiro_reais / cotação_dólar_atual
	fmt.Println("Convertendo a quantidade de dinheiro que foi informada em reais para dólar, você possui", dinheiro_dólares, "dólares.")
}
