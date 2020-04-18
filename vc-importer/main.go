package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Srno")
	f.SetCellValue("Sheet1", "B1", "LcoID")
	f.SetCellValue("Sheet1", "C1", "Vcno")
	f.SetCellValue("Sheet1", "D1", "Recharge_Amount")

	dat, err := ioutil.ReadFile("./number.txt")
	if err != nil {
		log.Fatalln(err)
	}

	split := strings.Split(string(dat), "\n")
	counts := 2
	for _, item := range split {
		if len(item) > 0 {
			nextsplit := strings.Split(item, "Pindi: ")
			if len(nextsplit) > 1 {
				mainNumberSplit := strings.Split(nextsplit[1], "+")
				number, _ := strconv.Atoi(mainNumberSplit[0])
				amount, _ := strconv.Atoi(mainNumberSplit[1])

				f.SetCellValue("Sheet1", "A"+strconv.Itoa(counts), counts-1)
				f.SetCellValue("Sheet1", "B"+strconv.Itoa(counts), 31045)
				f.SetCellValue("Sheet1", "C"+strconv.Itoa(counts), number)
				f.SetCellValue("Sheet1", "D"+strconv.Itoa(counts), amount)
			}
		}
		counts++
	}

	if err = f.SaveAs("Book1.xlsx"); err != nil {
		log.Fatalln(err)
	}
}
