package main

import (
	"bufio"
	"os"
	"strings"
)

func window() (string, []string, []string) {

	var paths, extensions []string
	reader := bufio.NewReader(os.Stdin)

	println("GO SEARCH")
	println("")
	println("Digite o termo a ser pesquisado: ")
	print("-> ")

	pattern, _ := reader.ReadString('\n')
	// convert CRLF to LF
	pattern = strings.Replace(pattern, "\n", "", -1)

	println("")
	println("Lista de caminhos: ")

	for {
		print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "" {
			break
		}
		paths = append(paths, text)
	}

	println("")
	println("Lista de extensões: ")

	for {
		print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "" {
			break
		}
		extensions = append(extensions, text)
	}

	return pattern, paths, extensions
}
