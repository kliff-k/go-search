package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	var pathList []string
	pathList = append(pathList, "E:\\Windows\\Games\\Touhou")
	pathList = append(pathList, "E:\\Windows\\Games\\RagnarokBattleOffline")

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
