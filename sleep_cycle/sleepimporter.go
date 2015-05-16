package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var fp *os.File
	var err error

	if len(os.Args) < 2 {
		fmt.Println("Insert file name or <C-C>.")
		fp = os.Stdin
	} else {
		fmt.Printf(">> read file %s\n", os.Args[1])
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, _, err := reader.ReadLine()
		fmt.Println(string(line))
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}

