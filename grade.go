//Autor: Ygor Pereira de Sá
//Data: 14/07/2021
//Descrição: Programa capaz de realizar a média de alunos usando desvio condicional para informar se o aluno foi aprovado ou reprovado.

package main

import (
	"fmt" //biblioteca de entada e saida de dados em Golang.
)

var nome string
var a1 float64
var a2 float64
var a3 float64
var media float64

func main() {

	fmt.Println("Qual o nome do aluno?")
	fmt.Scanln(&nome)
	fmt.Println("Qual é a nota da A1 do ", nome)
	fmt.Scanln(&a1)
	fmt.Println("Qual é a nota da A2 do aluno", nome)
	fmt.Scanln(&a2)
	fmt.Println("Qual é a nota da A2 do aluno", nome)
	fmt.Scanln(&a3)

	media = (a1 + a2 + a3) / 3

	if media < 6 {
		fmt.Println("Reprovou, estude mais na proxima")

	} else {
		fmt.Println("APROVADO, PARABÉNS!!!!")
	}

}
