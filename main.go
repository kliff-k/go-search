package main

import (
	"sync"
)

func main() {

	var waitGroup sync.WaitGroup

	pattern, paths, extensions := window()

	println("Resultados encontrados: ")
	println("")

	for _, path := range paths {

		waitGroup.Add(1)

		search(&waitGroup, pattern, path, extensions)

	}

	waitGroup.Wait()

}
