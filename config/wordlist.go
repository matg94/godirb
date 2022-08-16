package config

import (
	"bufio"
	"log"
	"os"

	"github.com/matg94/godirb/data"
)

func ReadAndCompileWordLists(queue *data.WordQueue, paths []string, words []string) {
	queue.AddList(words)
	for _, path := range paths {
		words, err := ReadWordlist(path)
		if err != nil {
			log.Fatal(err)
		}
		queue.AddList(words)
	}
}

func ReadWordlist(path string) ([]string, error) {
	readFile, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines, nil
}
