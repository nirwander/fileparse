package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	start := time.Now()
	fileBytes1, _ := ioutil.ReadFile(`D:\ry000067444`)
	fileBytes2, _ := ioutil.ReadFile(`D:\ry000067445`)

	//var bb bytes.Buffer

	/*

		ry000067444
		Len  1483 RBA 0
		Len  1974 RBA 16554

		ry000067445
		2018/07/04 16:19:54.877.341 FileHeader           Len  1805 RBA 0
		2018/06/25 22:17:57.212.518 Metadata             Len 93 RBA 1813
		2018/06/30 22:00:11.636.833 Metadata             Len 12061 RBA 1957
		2018/07/04 16:19:52.994.566 Insert               Len  1978 RBA 14088

	*/

	/*n, _ := bb.Write(fileBytes1[:1483])
	fmt.Printf("%d bytes buffered\n", n)

	n, _ = bb.Write(fileBytes2[1813:14088])
	fmt.Printf("%d bytes buffered\n", n)

	n, _ = bb.Write(fileBytes1[16654:])
	fmt.Printf("%d bytes buffered\n", n)

	err := ioutil.WriteFile(`D:\ry000067444n`, bb.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	*/
	fmt.Printf("%s\n\n", fileBytes1[:1815])
	fmt.Printf("%s\n", fileBytes2[:1815])
	//fmt.Printf("%x/n", bb.Bytes())
	fmt.Printf("\nDone in %s", time.Since(start))
}
