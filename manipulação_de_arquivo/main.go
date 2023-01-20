package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Cria arquivo
	f, err := os.Create("arquivos.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Hello, word!!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes", tamanho)
	f.Close()

	//leitura de arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//leitura de pouco em pouco abrindo arquivo

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	//cremover arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
