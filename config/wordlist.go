package config

import (
	"bufio"
	"log"
	"os"

	"github.com/matg94/godirb/data"
)

func ReadAndCompileWordLists(queue *data.WordQueue, paths []string, words []string, append []string) {
	if len(append) > 0 {
		queue.AddList(AppendWords(words, append))
	} else {
		queue.AddList(words)
	}
	for _, path := range paths {
		words, err := ReadWordlist(path)
		if err != nil {
			log.Fatal(err)
		}
		if len(append) > 0 {
			queue.AddList(AppendWords(words, append))
		} else {
			queue.AddList(words)
		}
	}
}

func AppendWords(words, appends []string) []string {
	var finalized_list []string
	for _, word := range words {
		for _, ext := range appends {
			finalized_list = append(finalized_list, word+ext)
		}
	}
	return finalized_list
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
