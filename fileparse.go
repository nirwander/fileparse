package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"
)

type repTable struct {
	srcOwner  []byte
	srcName   []byte
	tOwner    []byte
	tName     []byte
	extParams []byte
}

func main() {
	//fmt.Println("Hello World!")
	start := time.Now()
	processReplicat("C:\\Users\\wander\\go\\xfecr.txt")

	fmt.Printf("\n%s lines process\n", time.Since(start))
}

func processReplicat(fName string) {
	fileBytes, _ := ioutil.ReadFile(fName)

	//start = time.Now()
	lines := bytes.Split(fileBytes, []byte("\n"))
	//fmt.Printf("%s file split\n", time.Since(start))
	re := regexp.MustCompile("(?i)map[[:space:]]+([[:alnum:]_$]+)\\.([[:alnum:]_$\\?\\*\\-]+)[[:space:]]*,{0,1}[[:space:]]*target[[:space:]]+([[:alnum:]_$]+)\\.([[:alnum:]_$\\?\\*\\-]+)[[:space:]]*,{0,1}[[:space:]]*(.*);")

	repTables := make(map[string]repTable)
	var c2 int
	var c3 int
	for _, line := range lines {
		// Ищем предложения MAP OWNER.NAME TARGET OWNER.NAME [KEYCOLS (cols)] [params] ;
		//fmt.Printf("%d: %s", i, line)
		matches := re.FindSubmatch(line)
		c2++
		if len(matches) > 0 {
			//fmt.Printf("%q\n", matches)
			fmt.Printf("\t%s.%s >> %s.%s, tail: %s\n", matches[1], matches[2], matches[3], matches[4], matches[5])
			repTables[string(matches[3])+"."+string(matches[4])] = repTable{matches[1], matches[2], matches[3], matches[4], matches[5]}
			//str := string(matches[3]) + "." + string(matches[4])
			//fmt.Printf("\t%s\n", str)
			c3++
		}

		if bytes.Contains(line, []byte("Run Time Messages")) {
			break
		}
	}
	fmt.Printf("\n%d lines in file\n%d lines matched\n", c2, c3)
	fmt.Printf("%d tables in map", len(repTables))

}
