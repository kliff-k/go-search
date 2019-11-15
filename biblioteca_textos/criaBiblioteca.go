package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const DiferentWords int = 700000
const MaxWordRepetition int = 5
const NumberFolders int = 10
const WordsPerFile int = 100

var scanner *bufio.Scanner
var file *os.File

var listOfWords []string = make([]string, DiferentWords)
var wordCount int
var folderCount int = 1
var fileCount int = 1

func main() {
	openFile()
	numberOfFiles := DiferentWords * MaxWordRepetition / WordsPerFile
	fmt.Println("Number of Files: ", numberOfFiles)
	dir := createDirectory()
	createTextFile(500, dir)
}

func openFile() {
	file, err := os.Open("." + string(filepath.Separator) + "palavras_fonte.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		listOfWords[wordCount] = scanner.Text()
		wordCount++
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(listOfWords), func(i, j int) { listOfWords[i], listOfWords[j] = listOfWords[j], listOfWords[i] })

	fmt.Println("Finished creating array")
}

func getNextWord() string {
	if len(listOfWords) <= 0 {
		openFile()
	}
	newWord := ""
	for newWord == "" {
		randomIndex := rand.Intn(len(listOfWords))

		newWord = listOfWords[randomIndex]
		listOfWords = remove(listOfWords, randomIndex)
	}
	//fmt.Println(newWord)
	return newWord
}

func createTextFile(numberOfWords int, folderName string) {
	slash := string(filepath.Separator)
	dir := folderName + slash + "livro-" + strconv.Itoa(fileCount) + ".txt"
	fmt.Println(dir)
	file, err := os.Create(dir)
	defer file.Close()
	check(err)

	fileCount++

	w := bufio.NewWriter(file)
	for numberOfWords > 0 {
		w.WriteString(getNextWord() + "\n")
		numberOfWords--
	}

	w.Flush()
}

func createDirectory() string {
	folderName := "pasta" + strconv.Itoa(folderCount)
	folderCount++
	dir := "." + string(filepath.Separator) + folderName
	os.Mkdir(dir, 0777)
	//fmt.Println(dir)
	return dir
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
