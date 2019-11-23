package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// Função principal de busca
func search(waitGroup *sync.WaitGroup, pattern string, root string, extensions []string, grep bool) {

	// Constroi um objeto de expressão regular com base na entrada do usuário
	regEx, errRegEx := regexp.Compile("(?i)" + pattern)
	if errRegEx != nil {
		log.Fatal(errRegEx)
	}

	if root == "" {
		root = "."
	}

	// Caminha toda a árvore de diretórios com base no caminho fornecido pelo usuário
	errWalk := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		// O sistema retorna somente arquivos, então ignora diretórios
		if info.IsDir() {
			return nil
		}

		// Se foi fornecido alguma extensão, verifica se o arquivo pertence a este tipo
		var extensionValid = false
		if len(extensions) > 0 {
			for _, ext := range extensions {
				if strings.HasSuffix(info.Name(), ext) {
					extensionValid = true
				}
			}
		} else {
			extensionValid = true
		}

		if !extensionValid {
			return nil
		}

		// Verifica se o nome do arquivo contem o termo fornecido
		if regEx.MatchString(info.Name()) {
			println(info.Name())
			return nil
		}

		// Verifica se o conteúdo do arquivo contém alguma menção ao termo fornecido
		if grep {
			file, errFile := ioutil.ReadFile(path)
			if errFile != nil {
				log.Fatal(errFile)
			}
			fileContent := string(file)
			if regEx.MatchString(fileContent) {
				println(info.Name())
				return nil
			}
		}

		return nil
	})

	if errWalk != nil {
		log.Fatal(errWalk)
	}

	// Indica que a rotina terminou
	waitGroup.Done()

}
