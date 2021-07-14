//Autor: Ygor Pereira de Sá
//Data: 14/07/2021
//Descrição: Código irá calcular a média escolar dos aulunos, criado para entender os operadores matemáticos do Golang

package main

import "fmt"

var name string
var a1 float64
var a2 float64
var a3 float64
var media float64

func main() {

	for i := 0; i < 10; i++ {

		fmt.Println("Qual o nome do aluno", i+1, "?\n")
		fmt.Scanln(&name)

		fmt.Println("Qual a nota da A1 do aluno ", i+1, "?\n")
		fmt.Scanln(&a1)

		fmt.Print("Qual a nota da A2 do aluno ", i+1, "?\n")
		fmt.Scanln(&a2)

		fmt.Println("Qual a nota da A3 do aluno ", i+1, "?\n")
		fmt.Scanln(&a3)

		media = (a1 + a2 + a3) / 3

		fmt.Println("A média do aluno", i+1, "é: ", media)

	}

}
