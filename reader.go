package main

import (
	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func Read() {
	path := "./files/ASTRUM SOH 01 12 2021.xlsx"
	//path := "./files/De Wet Stock List Template.csv"
	paths := []string{path}

	for value := range paths {
		ext := filepath.Ext(paths[value])
		switch ext {
		case ".csv":
			fmt.Println("reading CSV")
			readCSV(path)
		case ".xlsx":
			fmt.Println("reading XLSX")
			readXLSX(path)
		}
	}
}

func readCSV(p string) {
	f, err := os.Open(p)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {
		records, err := r.ReadAll()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		for _, record := range records {
			isEmpty := empty(record)
			if isEmpty {
				fmt.Printf("%s\n", record)
			}
		}
	}
}

func empty(record []string) bool {
	for i := range record {
		entry := record[i]

		isNotUTF8 := isNotUTF8(entry)
		if isNotUTF8 || entry == "" {
			return false
		}
	}

	return true
}

func readXLSX(p string) {
	f, err := excelize.OpenFile(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		//for _, colCell := range row {
		//	fmt.Print(colCell, "\t")
		//}
		fmt.Println(row)
	}
}

func isNotUTF8(s string) bool {
	re, err := regexp.Compile(`[^\x00-\x7F]+`)
	if err != nil {
		log.Fatal(err)
	}
	return re.MatchString(s)
}
