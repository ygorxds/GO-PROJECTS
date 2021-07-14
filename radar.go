//Autor: Ygor Pereira de Sá
//Data: 14/07/2021
//Descrição: Código irá calcular a velocidade média de um carro, considerei como o calculo de um radar.
package main

import (
	"fmt"
)

var t float64
var vm float64

func main() {

	s := 1000.0

	fmt.Println("Insira o tempo em segundos:")
	fmt.Scanln(&t)
	vm = s / t

	if vm > 22.22 {
		fmt.Println("Você foi multado safado, ande despacito na próxima")
	} else {

	}
	fmt.Println("A velocidade média é", vm, "M/S")

}
