package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	var dateStr string
	flag.StringVar(&dateStr, "date", "2021-01-03", "A date format string, do not add time, it will added automatically")
	flag.Parse()
	dateStr = dateStr + "T00:04:05.000Z"
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		log.Fatalln(err)
	}
	_, wkNum := t.ISOWeek()
	//timeNow := time.Now()
	//_, wkNum := timeNow.ISOWeek()

	fmt.Println("Week number for the date", t.Format("2006-01-02"), "is", wkNum)
	//fmt.Println(t.Format("2006-01-02 15:04:05"))
	var x string
	fmt.Scanln(&x)
}
