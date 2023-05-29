package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ParseCSV(filePath string) ([]Table, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ファイルを開けません: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	rawCsvData, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("CSVを読み込み中にエラーが発生しました: %v", err)
	}

	if len(rawCsvData) < 2 {
		return nil, fmt.Errorf("CSVファイルには最低でも2行が必要です")
	}

	return ParseCSVData(rawCsvData), nil
}

func ParseCSVData(rawCsvData [][]string) []Table {
	tables := make(map[string]*Table)

	for _, record := range rawCsvData[1:] {
		tableName := record[0]
		if _, ok := tables[tableName]; !ok {
			tables[tableName] = &Table{Name: tableName, Columns: []Column{}}
		}

		column := Column{
			Name:             record[1],
			Type:             record[2],
			IsPrimaryKey:     ParseBool(record[3]),
			IsNotNull:        ParseBool(record[4]),
			IsUnique:         ParseBool(record[5]),
			ForeignKeyTable:  record[6],
			ForeignKeyColumn: record[7],
			Check:            record[8],
			Comment:          record[9],
		}

		tables[tableName].Columns = append(tables[tableName].Columns, column)
	}

	var tableList []Table
	for _, table := range tables {
		tableList = append(tableList, *table)
	}
	return tableList
}
