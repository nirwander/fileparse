package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	fileBytes1, _ := ioutil.ReadFile(`D:\ry000067444`)
	fileBytes2, _ := ioutil.ReadFile(`D:\ry000067445`)

	var bb bytes.Buffer
	//bb = bytes.Buffer(fileBytes1[:1483])
	n, _ := bb.Write(fileBytes1[:1483])
	fmt.Printf("%d bytes buffered\n", n)
	fmt.Println(bb.Bytes())
}
