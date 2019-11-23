package main

import (
	"bufio"
	"os"
	"strings"
)

// Função da interface com o usuário
func window() (string, []string, []string) {

	var paths, extensions []string
	reader := bufio.NewReader(os.Stdin)

	// Exibe a tela inicial
	println("GO SEARCH")
	println("")
	println("Digite o termo a ser pesquisado: ")
	print("-> ")

	// Recupera o termo a ser buscado pelo programa
	pattern, _ := reader.ReadString('\n')
	pattern = strings.Replace(pattern, "\n", "", -1)

	// Exibe o prompt de caminhos
	println("")
	println("Lista de caminhos: ")

	// Captura caminhos até que uma linha vazia seja fornecida
	for {
		print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "" {
			break
		}
		paths = append(paths, text)
	}

	// Exibe o prombt de extensões
	println("")
	println("Lista de extensões: ")

	// Captura extensões até que uma linha vazia seja fornecida
	for {
		print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "" {
			break
		}
		extensions = append(extensions, text)
	}

	// Retorna a informações fornecidas pelo usuário para o main utilizar na busca
	return pattern, paths, extensions
}
