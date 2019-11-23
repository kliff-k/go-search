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

func search(waitGroup *sync.WaitGroup, pattern string, root string, extensions []string) {
	regEx, errRegEx := regexp.Compile("(?i)" + pattern)
	if errRegEx != nil {
		log.Fatal(errRegEx)
	}

	errWalk := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		var extensionValid = false
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			return nil
		}

		if len(extensions) > 0 {
			for _, ext := range extensions {
				if strings.HasSuffix(info.Name(), ext) {
					extensionValid = true
				}
			}
		}

		if !extensionValid {
			return nil
		}

		file, err := ioutil.ReadFile(root + info.Name())
		fileContent := string(file)

		if regEx.MatchString(info.Name()) || regEx.MatchString(fileContent) {
			println(info.Name())
		}

		return nil
	})

	if errWalk != nil {
		log.Fatal(errWalk)
	}

	waitGroup.Done()

}
