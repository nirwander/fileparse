package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	start := time.Now()

	fileBytes, _ := ioutil.ReadFile("C:\\Users\\wander\\go\\xfecr.txt")

	lines := bytes.Split(fileBytes, []byte("\n"))
	re := regexp.MustCompile("(?i)map[[:space:]]+([[:alnum:]_$]+)\\.([[:alnum:]_$\\?\\*\\-]+)[[:space:]]*,{0,1}[[:space:]]*target[[:space:]]+([[:alnum:]_$]+)\\.([[:alnum:]_$\\?\\*\\-]+)[[:space:]]*,{0,1}[[:space:]]*(.*);")

	for i, line := range lines {
		// Ищем предложения MAP OWNER.NAME TARGET OWNER.NAME [KEYCOLS (cols)] [params] ;
		fmt.Printf("%d: %s", i, line)
		matches := re.FindSubmatch(line)
		if len(matches) > 0 {
			//fmt.Printf("%q\n", matches)
			fmt.Printf("\t%s.%s >> %s.%s, tail: %s\n", matches[1], matches[2], matches[3], matches[4], matches[5])
		}

	}
	fmt.Println(time.Since(start))

}
