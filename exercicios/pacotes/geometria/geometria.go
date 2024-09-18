/*
2. Estrutura de múltiplos pacotes:
  - Crie dois pacotes: `geometria` e `trigonometria`.
  - Em `geometria`, implemente funções para calcular área e perímetro de figuras geométricas como círculo, quadrado e retângulo.
  - Em `trigonometria`, crie funções que calculam seno, cosseno e tangente de um ângulo.
  - No pacote principal, use funções de ambos os pacotes para realizar cálculos.
*/

package geometria

import "ex1/matematica"

func Area[T matematica.Number](lado1 T, lado2 T) T {
	return lado1 * lado2
}
