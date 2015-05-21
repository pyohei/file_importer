package main

import (
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
)

const LOG_FILE = "sleepimporter.log"

func main() {

	wf, _ := os.OpenFile(LOG_FILE, os.O_WRONLY|os.O_CREATE, 0755)
	log.SetOutput(wf)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("START!")
	if len(os.Args) != 2 {
		log.Println("You should hava argument with filename")
		os.Exit(1)
	}
	//fmt.Printf(">> read file %s\n", os.Args[1])
	// if possible, return with iterable
	filelines := fileReader(os.Args[1])
	//fmt.Println(filelines)

	/*
		// それができたら、ひとまずこのbatchは終わりとする。
		// 統計結果を出すのはまた別。
		// SQLの見直しも必要
	*/

	for i, line := range filelines {
		if i == 0 {
			continue
		}
		err := insertRecord(line)
		if err != nil {
			log.Printf("Line: %v Error!!", i)
			panic(err)
		}
		log.Printf("Line: %v Succcess!!", i)
	}
	log.Println("FINISH")
	wf.Close()
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
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return bbb
}

func insertRecord(rec []string) error {
	conn, err := sql.Open(
		"mysql",
		"developer:developer@tcp(192.168.0.90:3306)/sleep_cycle?charset=utf8")
	if err != nil {
		return err
	}
	stmt, err := conn.Prepare(
		"INSERT sleep SET sleep_from=?, sleep_to=?, sleep_len=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(rec[0], rec[1], rec[3])
	_ = res
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
