package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	csvPath := flag.String("csv", "", "CSVファイル名")
	ddlPath := flag.String("output", "", "DDLファイル名")
	dirName := flag.String("dir", "", "複数のCSVファイルが格納されているディレクトリパス")
	flag.Parse()

	if *csvPath != "" && *dirName != "" {
		fmt.Println("Error: CSVファイル名、またはディレクトリパスが指定されていません。")
		os.Exit(1)
	}

	if *csvPath != "" {
		if *ddlPath == "" {
			*ddlPath = strings.TrimSuffix(*csvPath, filepath.Ext(*csvPath)) + ".sql"
		}
		processFile(*csvPath, *ddlPath)
	} else if *dirName != "" {
		files, err := ioutil.ReadDir(*dirName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, f := range files {
			if filepath.Ext(f.Name()) == ".csv" {
				csvFile := filepath.Join(*dirName, f.Name())
				outputFile := strings.TrimSuffix(csvFile, filepath.Ext(csvFile)) + ".sql"
				if !processFile(csvFile, outputFile) {
					fmt.Println("未対応のCSVファイル: ", csvFile)
				}
			}
		}
	} else {
		fmt.Println("Error: CSVファイル名、またはディレクトリパスが指定されていません。")
		os.Exit(1)
	}
}

func processFile(csvPath string, ddlPath string) bool {
	tables, err := ParseCSV(csvPath)
	if err != nil {
		fmt.Println(err)
		return false
	}

	ddl, err := GenerateDDL(tables)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if err := WriteToFile(ddlPath, ddl); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
