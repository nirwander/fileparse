package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	fmt.Println("Hello World!")

	fileBytes, _ := ioutil.ReadFile("C:\\Users\\wander\\go\\xfecr.txt")

	lines := bytes.Split(fileBytes, []byte("\n"))
	re := regexp.MustCompile("(?i)map[[:space:]]+([a-z]+)\\.([a-z]+)")

	for i, line := range lines {
		// Ищем предложения MAP OWNER.NAME TARGET OWNER.NAME [KEYCOLS (cols)] [params] ;
		fmt.Printf("%d: %s", i, line)

	}

}
