package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	csvFile := flag.String("csv", "", "CSVファイルへのパス.")
	outputName := flag.String("output", "", "DDLの出力ファイル名。指定しない場合、CSVファイルの名前が使用されます。")
	dir := flag.String("dir", "", "CSVファイルが含まれるディレクトリのパス.")
	singleFile := flag.Bool("single", false, "-dirフラグと共に使用。全てのDDLを一つのファイルにまとめます。")
	flag.Parse()

	if *csvFile != "" && *dir != "" {
		fmt.Println("フラグは-csvか-dirのどちらか一つだけを指定してください。")
		return
	}

	if *csvFile == "" && *dir == "" {
		fmt.Println("フラグ-csvか-dirのどちらか一つを指定してください。")
		return
	}

	if *singleFile && *dir == "" {
		fmt.Println("フラグ-singleは-dirフラグと共に使用してください。")
		return
	}

	if *csvFile != "" {
		ddl, err := ProcessCSVFile(*csvFile)
		if err != nil {
			fmt.Println("エラー：", err)
		}
		if err := WriteToFile(*outputName+".sql", ddl); err != nil {
			fmt.Println("エラー：", err)
		}
		return
	}

	if *dir != "" {
		csvFiles, err := GetFilesFromDirectory(*dir, ".csv")
		if err != nil {
			fmt.Println("エラー：", err)
			return
		}

		var allDDLs string
		for _, file := range csvFiles {
			filename := filepath.Base(file)
			filename = filename[:len(filename)-len(filepath.Ext(filename))]
			if *singleFile {
				ddl, err := ProcessCSVFile(file)
				if err != nil {
					fmt.Printf("ファイルの処理エラー：%s。 エラー：%s\n", file, err)
					continue
				}
				allDDLs += ddl + "\n"
			} else {
				ddl, err := ProcessCSVFile(file)
				if err != nil {
					fmt.Printf("ファイルの処理エラー：%s。 エラー：%s\n", file, err)
				}
				if err := WriteToFile(filename+".sql", ddl); err != nil {
					fmt.Println("エラー：", err)
				}
			}
		}

		if *singleFile {
			outputFileName := "output.sql"
			if *outputName != "" {
				outputFileName = *outputName + ".sql"
			}
			if err := WriteToFile(outputFileName, allDDLs); err != nil {
				fmt.Println("エラー：", err)
			}
		}
	}
}
