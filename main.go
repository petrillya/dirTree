package main

import (
	"fmt"
	"os"
)

func listDir(dir string, count int) error {
	var haveFolder bool = false
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, entrie := range entries {
		if entrie.IsDir() {
			haveFolder = true
			break
		}
	}

	if haveFolder {
		count += 1
	} else {
		count -= 1
	}

	for _, entrie := range entries {
		if entrie.IsDir() {
			for i := 1; i < count; i++ {
				fmt.Print("\t")
			}

			fmt.Println("├───" + entrie.Name())

			haveFolder = false
			listDir(dir+"/"+entrie.Name(), count)
		}
	}

	return nil
}

func main() {
	listDir("./", 0)
}
