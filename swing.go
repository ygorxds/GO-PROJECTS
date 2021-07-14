//Autor: Ygor Pereira de Sá
//Data: 14/07/2021
//Descrição: Código irá exibir a categoria de um competidor de acordo com a idade com o uso de laço de repetição.

package main

import (
	"fmt"
)

var age int

func main() {

	fmt.Println("Qual a idade do competidor?")
	fmt.Scanln(&age)
	if age >= 5 && age <= 7 {
		fmt.Println("Categoria Infantil A")
	} else if age >= 8 && age <= 10 {
		fmt.Println("Categoria Infantil B")
	} else if age >= 11 && age <= 13 {
		fmt.Println("Juvenil A")
	} else if age >= 14 && age <= 17 {
		fmt.Println("Juvenil B")
	} else if age > 18 {
		fmt.Println("senior")
	} else {
		fmt.Println("Failing grade")
	}

}
