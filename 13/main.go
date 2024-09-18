package main

import (
	"fmt"
	"primeiro_pacote/math"
)

func main() {
	soma := math.Soma(1, 2)
	fmt.Println("Pacotes")

	fmt.Println(soma)

	//fmt.Println(math.a) Não compila, pois "variaveis minúsculas" só são vistas dentro do pacote

	fmt.Println(math.A) //Variáveis maiúsculas podem ser usadas fora do pacote

	// O mesmo vale para structs com inicial maiúscula ou minúscula
	// e o mesmo para os campos da struct que estão iniciados
	// com letra maiúscula ou minúscula
}
