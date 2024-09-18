/*
2. Estrutura de múltiplos pacotes:
  - Crie dois pacotes: `geometria` e `trigonometria`.
  - Em `geometria`, implemente funções para calcular área e perímetro de figuras geométricas como círculo, quadrado e retângulo.
  - Em `trigonometria`, crie funções que calculam seno, cosseno e tangente de um ângulo.
  - No pacote principal, use funções de ambos os pacotes para realizar cálculos.
*/

package trigonometria

import "ex1/matematica"

func Seno[T matematica.Number](altura T, percurso T) T {
	return altura / percurso
}
