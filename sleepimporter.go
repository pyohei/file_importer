package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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
	filelines := fileReader(os.Args[1])

	for i, line := range filelines {
		fmt.Println(i)
		if i == 0 {
			continue
		}
		hasRec := hasRecord(line)
		if hasRec {
			log.Printf("Line: %v has already record!!", i)
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
	bbb := make([][]string, 0, 1000)
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

func hasRecord(rec []string) bool {
	var sCount int

	conn, err := sql.Open(
		"mysql",
		"developer:developer@tcp(192.168.0.90:3306)/sleep_cycle?charset=utf8")
	row := conn.QueryRow(
		"SELECT count(*) as sCount FROM sleep WHERE sleep_from = ? and sleep_to =?",
		rec[0],
		rec[1])
	err = row.Scan(&sCount)
	conn.Close()
	if err != nil {
		panic(err)
	}
	if sCount > 0 {
		return true
	}
	return false
}

func insertRecord(rec []string) error {
	var cRate string
	var sMinute int

	conn, err := sql.Open(
		"mysql",
		"developer:developer@tcp(192.168.0.90:3306)/sleep_cycle?charset=utf8")
	if err != nil {
		return err
	}
	stmt, err := conn.Prepare(
		"INSERT sleep SET sleep_from=?, sleep_to=?, confort_rate=?, sleep_minute=?, sleep_feeling=?, memo=?, pulsation=?, walk_count=?, regist_time=?")
	if err != nil {
		return err
	}
	// Confort Rate
	if rec[2] == "" {
		cRate = "0"
	} else {
		cRate = strings.Replace(rec[2], "%", "", -1)
	}
	// Sleep minute
	if rec[3] == "" {
		sMinute = 0
	} else {
		sTimes := strings.Split(rec[3], ":")
		sHour, _ := strconv.Atoi(sTimes[0])
		sMin, _ := strconv.Atoi(sTimes[1])
		sMinute = sHour*60 + sMin
	}
	now := time.Now()
	feelNum := convertFeeling(rec[4])
	res, err := stmt.Exec(
		rec[0], rec[1], cRate, sMinute,
		feelNum, rec[5], convertNull(rec[6]),
		rec[7], now.Format("2006-01-02 15:04:05"))
	_ = res
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func convertNull(rec string) string {
	if rec == "" {
		return "null"
	}
	return rec
}

func convertFeeling(feeling string) int {
	var feelNum int
	if feeling == ":(" {
		feelNum = 1
	} else if feeling == ":|" {
		feelNum = 2
	} else if feeling == ":)" {
		feelNum = 3
	} else {
		feelNum = 0
	}
	return feelNum
}
