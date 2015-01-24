package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writeToFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("test")
}

func readFile(name string) (content string) {
	bs, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}

	content = string(bs)

	return content
}

func main() {
	const fileName = "foo.txt"
	writeToFile(fileName)
	fileStuff := readFile(fileName)
	fmt.Println(fileStuff)
}
