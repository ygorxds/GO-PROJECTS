//Autor: Ygor Pereira de Sá
//Data: 14/07/2021
//Descrição: Código irá calcular a tabuada de acordo com a decisão do usuário.

package main

import (
	"fmt"
)

var ops int
var num int
var soma int
var mlt int
var div int
var sub int

func main() {

	fmt.Println("Qual o tipo de operação você quer realizar\n")
	fmt.Println(" Digite 1 para (+)")
	fmt.Println(" Digite 2 para (-)")
	fmt.Println(" Digite 3 para (*)")
	fmt.Println(" Digite 4 para (/)")
	fmt.Scanln(&ops)
	fmt.Println("Qual a tabuada que deseja saber?")
	fmt.Scanln(&num)
	if ops == 1 {
		for i := 0; i < 10; i++ {
			soma = (i + 1) + num
			fmt.Println(num, "+", i+1, "=", soma, "\n")
			fmt.Print("")

		}

	} else if ops == 2 {
		for i := 0; i < 10; i++ {
			sub = (i + 1) - num
			fmt.Println(num, "-", i+1, "=", sub, "\n")
			fmt.Print("")

		}
	} else if ops == 3 {
		for i := 0; i < 10; i++ {
			mlt = (i + 1) * num
			fmt.Println(num, "X", i+1, "=", mlt, "\n")
			fmt.Print("")
		}

	} else if ops == 4 {
		for i := 0; i < 10; i++ {
			div = num / (i + 1)
			fmt.Println(num, "/", i+1, "=", div, "\n")
			fmt.Print("")

		}
	} else {
		fmt.Println("INSIRA UM SINAL VÁLIDO CACHORRO!!!")
	}
}
