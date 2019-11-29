package main

import (
	"os"
	"sync"
)

// Função principal do sistema.
func main() {

	// Inicia o waitgroup utilizado pelas rotinas
	var waitGroup sync.WaitGroup
	var grep bool = false

	// Recupera os argumentos da chamada ao programa
	args := os.Args

	// Ativa o modo GREP
	if len(args) > 1 {
		if args[1] == "-grep" {
			grep = true
		}
	}

	// Recupera os parametros passados pelo usuário
	pattern, paths, extensions := window()

	// Inicia a lista de resultados
	println("Resultados encontrados: ")
	println("")

	// Laço principal. Gera uma rotina para cada caminho indicado pelo usuário
	for _, path := range paths {

		// Adiciona ao waitgroup
		waitGroup.Add(1)

		// Executa a busca por nome e conteúdo
		go search(&waitGroup, pattern, path, extensions, grep)

	}

	// Aguarda o fim de todas as rotinas antes de finalizar o programa
	waitGroup.Wait()

}
