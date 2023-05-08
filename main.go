package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	csvPath := flag.String("csv", "", "CSVファイル名")
	ddlPath := flag.String("output", "", "DDLファイル名")
	flag.Parse()

	if *csvPath == "" {
		fmt.Println("Error: CSVファイル名が指定されていません。")
		os.Exit(1)
	}

	if *ddlPath == "" {
		*ddlPath = filepath.Base(*csvPath)
		ext := filepath.Ext(*ddlPath)
		*ddlPath = (*ddlPath)[:len(*ddlPath)-len(ext)] + ".sql"
	}

	tables, err := ParseCSV(*csvPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	ddl, err := GenerateDDL(tables)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(*ddlPath, []byte(ddl), 0644)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("DDLファイル: %s\n", *ddlPath)
}
