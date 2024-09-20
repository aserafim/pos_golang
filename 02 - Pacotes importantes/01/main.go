package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Manipulando arquivos")

	// Criação/Escrita de arquivo
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo criado."))

	if err != nil {
		panic(err)
	}
	fmt.Printf("Tamanho do arquivo: %d bytes.\n", tamanho)
	f.Close()

	// Leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// o casting para string
	// é necessário pois os arquivos
	// armazenam bytes e não chars/strings
	fmt.Println(string(arquivo))

	// Como ler arquivos muito grandes??
	// por exemplo, um arquivo que tenha
	// um tamanho maior do que a memória
	// disponível?

	// leitura parcial e contínua
	file, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Removendo arquivo
	erro := os.Remove("arquivo.txt")
	if err != nil {
		panic(erro)
	}

}
