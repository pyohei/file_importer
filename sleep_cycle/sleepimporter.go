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
	// 上記のelseをなくし、ここからファイルの読み込み->
	// 一覧の取得->DBへの追加をするようにする
	// それができたら、ひとまずこのbatchは終わりとする。
	// 統計結果を出すのはまた別。
	// SQLの見直しも必要
	fmt.Println(fileReader("sleepdata.csv"))

	fmt.Println("start")
	cnn, err := sql.Open(
		"mysql",
		"developer:developer@tcp(192.168.0.90:3306)/sleep_cycle?charset=utf8")
	stmt, err := cnn.Prepare(
		"INSERT sleep SET sleep_from=?, sleep_to=?, sleep_len=?")
	res, err := stmt.Exec("time", "time", "time")
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	/*
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
	*/
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

func fileReader(filepath string) [][]string {
	var fp *os.File
	var err error
	fp, err = os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
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
	return bbb
}
