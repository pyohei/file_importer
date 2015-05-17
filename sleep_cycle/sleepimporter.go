package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	//cnn := sql.Open("mysql", "developer:developer@tcp(192.168.0.90:3306)/hatena")
	cnn, err := sql.Open("mysql", "developer:developer@tcp(192.168.0.90:3306)/hatena?charset=utf8")
	rows, err := cnn.Query("select user_name from users;")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user_name string
		if err := rows.Scan(&user_name); err != nil {
			fmt.Println(err)
		}
		fmt.Println(user_name)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	bbb := make([][]string, 0, 200)
	scanner := bufio.NewScanner(fp)
	for i := 0; scanner.Scan(); i++ {
		r := strings.NewReplacer("\n", "")
		my_line := strings.Split(r.Replace(scanner.Text()), ";")
		//aaa := make([]string, 0, 200)
		//aaa = append(aaa, my_line[0])
		//aaa = append(aaa, my_line[1])
		bbb = append(bbb, my_line)
		//fmt.Print(aaa)
		//	fmt.Println(len(my_line))
	}
	fmt.Println(bbb)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
