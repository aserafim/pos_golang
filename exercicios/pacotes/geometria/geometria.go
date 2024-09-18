package geometria

import "ex1/matematica"

func Area[T matematica.Number](lado1 T, lado2 T) T {
	return lado1 * lado2
}
