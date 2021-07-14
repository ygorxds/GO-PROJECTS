//Autor: Ygor Pereira de Sá
//Data: 14/07/2021
//Descrição: Código irá calcular a média escolar dos aulunos, usando Arrey (vetor e laço de repetição)

package main

import (
	"fmt"
)

var nome [5]string
var a1 [5]float64
var a2 [5]float64
var a3 [5]float64
var media [5]float64

func main() {
	fmt.Println("Bem vindo ao sistema de notas da escola IFCMNV (INSTITUTO FILHO CHORA E MÃE NÃO VÊ\n")
	for i := 0; i < 5; i++ {
		fmt.Println("Insire o nome do aluno", i+1, ":\n")
		fmt.Scanln(&nome[i])
		fmt.Println("Insire a Nota da A1 do aluno", i+1, ":\n")
		fmt.Scanln(&a1[i])
		fmt.Println("Insire a Nota da A2 do aluno", i+1, ":\n")
		fmt.Scanln(&a2[i])
		fmt.Println("Insire a Nota da A3 do aluno", i+1, ":\n")
		fmt.Scanln(&a3[i])

		media[i] = (a1[i] + a2[i] + a3[i]) / 3

		if media[i] > 5.9 {

			fmt.Print(media[i], " TÁ APROVADO MEU CONSAGRETED\n")

		} else {
			fmt.Print(media[i], " REPROVADO, TENTE NOVAMENTE ANO QUE VEM, FELIZ NATAL!!!\n")
		}

	}

}
