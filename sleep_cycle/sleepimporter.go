package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	fmt.Println("start")
	//	lines := make([]string, 0)

	//	my_lines := make([]string, 120)
	scanner := bufio.NewScanner(fp)
	for i := 0; scanner.Scan(); i++ {
		r := strings.NewReplacer("\n", "")
		my_line := strings.Split(r.Replace(scanner.Text()), ";")
		fmt.Println("----")
		fmt.Println(my_line)
		fmt.Println(len(my_line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
