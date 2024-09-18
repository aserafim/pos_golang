package trigonometria

import "ex1/matematica"

func Seno[T matematica.Number](altura T, percurso T) T {
	return altura / percurso
}
