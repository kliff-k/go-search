package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	var pathList []string

	// Esta parte deve vir da interface gráfica.
	pathList = append(pathList, "/home/<user>/comics")
	pathList = append(pathList, "/mnt/<hd-externo>/hqs")

	// As chamadas devem se adaptar para a biblioteca / busca

	// Aqui dentro, cada path informado acima deve inicializar uma goroutine (Pegar de um dos exemplos)
	for _, path := range pathList {

		files, err := ioutil.ReadDir(path)

		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}

}
