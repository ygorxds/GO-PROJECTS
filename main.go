//Programa de fixa de usuário.
//Autor: Ygor Pereira de Sá
//08/07/2021
// criado para entender os tipos de váriáveis do golang.

package main

import (
	"fmt" //biblioteca de entada e saida de dados em Golang.
	"math"
)

var nome string
var idade int
var peso float64
var altura float64
var imc float64

func main() {

	fmt.Println("Qual o seu nome?\n")
	fmt.Scanln(&nome)
	fmt.Println("Qual o seu peso?\n")
	fmt.Scanln(&peso)
	fmt.Println("Qual a sua altura?\n")
	fmt.Scanln(&altura)
	fmt.Println("Qual a sua idade?\n")
	fmt.Scanln(&idade)
	imc = (peso / math.Pow(altura, 2))
	fmt.Println("Olá", nome, "O seu IMC é :", imc, "Sua altura é: ", altura, "Seu peso é: ", peso)

}
