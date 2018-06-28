package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello World!")

	fileBytes, _ := ioutil.ReadFile("C:\\Users\\wander\\go\\xfecr.txt")

	lines := bytes.Split(fileBytes, []byte("\n"))

	for i, line := range lines {
		fmt.Printf("%d: %s", i, line)
	}

}
