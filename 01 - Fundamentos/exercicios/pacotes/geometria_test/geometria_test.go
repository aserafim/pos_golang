/*
5. Pacote utilitário:
  - Crie um pacote utilitário chamado `stringsutil` que tenha funções para manipulação de strings, como inverter uma string, contar vogais e verificar se uma string é palíndroma.
  - Escreva testes para essas funções no pacote `stringsutil_test`.
*/

package geometriatest

import (
	"ex1/geometria"
	"fmt"
	"testing"
)

func Test_Area(t *testing.T) {
	lado1 := 40
	lado2 := 20

	fmt.Println(geometria.Area(lado1, lado2))
}
