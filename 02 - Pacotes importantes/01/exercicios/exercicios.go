// Exercícios para praticar:

/*
1. Abrir arquivo existente:
Modifique o código para abrir um arquivo já existente ("arquivo.txt"),
escreva mais dados no final do arquivo sem sobrescrevê-lo, e depois leia o arquivo inteiro para verificar os novos dados.

2. Contar linhas em um arquivo:
Crie um programa que conte e imprima o número de linhas em um arquivo de texto.
Use um buffer para ler o arquivo linha por linha.

3. Dividir a leitura em pedaços:
Modifique o programa para ler um arquivo maior, por exemplo 50 bytes de cada vez, ao invés de 10.

4. Escrever e ler dados estruturados:
Crie um arquivo para salvar informações estruturadas, como nome e idade de pessoas.
Em seguida, leia essas informações do arquivo e exiba no formato adequado.

5. Cópia de arquivos:
Implemente um programa que copie o conteúdo de um arquivo para outro arquivo.
Certifique-se de que funcione com arquivos grandes.

6. Manipular arquivos CSV:
Escreva um programa para ler um arquivo CSV e imprimir cada linha de forma formatada.
*/
package main

import (
	//	"bufio"

	"fmt"
	"os"
)

func main() {

	fmt.Println("Exercícios")

	/*
		1. Abrir arquivo existente:
		Modifique o código para abrir um arquivo já existente ("arquivo.txt"),
		escreva mais dados no final do arquivo sem sobrescrevê-lo, e depois leia o arquivo inteiro para verificar os novos dados.
	*/

	file, err := os.OpenFile("/home/aserafim/dev-repos/go-env/pos_golang/02 - Pacotes importantes/01/exercicios/teste.txt", os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Fprintln(file, "Escrevendo no arquivo já existente.")

	//_, err = file.WriteString("Adicionando informação no arquivo já existente.")
	file.Sync()
	defer file.Close()

	/*
		2. Contar linhas em um arquivo:
		Crie um programa que conte e imprima o número de linhas em um arquivo de texto.
		Use um buffer para ler o arquivo linha por linha.
	*/

	file, err = os.Open("/home/aserafim/dev-repos/go-env/pos_golang/02 - Pacotes importantes/01/exercicios/teste2.txt")
	//	scanner := bufio.NewReader(file)
	//scanner.Split(bufio.ScanLines())

	//	bufio.ScanLines(file)

}
