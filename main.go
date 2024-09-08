package main

import (
	"fmt"
	"os"
)

func listDir(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, entrie := range entries {
		if entrie.IsDir() {
			fmt.Println(entrie.Name())

			listDir(dir + "/" + entrie.Name())
		}
	}

	return nil
}

func main() {
	listDir("./")
}
