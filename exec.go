package main

import (
        "fmt"
        "log"
        "os/exec"
        "bytes"
)

func main() {
        fmt.Println("Hello World!")

        execCmd("/home/oracle/product/ggate12c/ggsci")
}

func execCmd (bin string) {
        var out bytes.Buffer

        cmd := exec.Command(bin)
        //cmd.Stdin = bytes.NewBuffer([]byte("info all"))
        cmd.Stdin = bytes.NewBuffer([]byte("view report XFECR"))
        cmd.Stdout = &out

        err := cmd.Run()

        if err != nil {
                log.Fatal(err)
        }


        lines := bytes.Split(out.Bytes(), []byte("\n"))

        for i, line := range lines {
                fmt.Printf("%d: %s\n", i, line)
        }

        //fmt.Printf("Output:\n%s\n", out.Bytes() )

}
