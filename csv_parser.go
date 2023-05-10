package main

import (
	"encoding/csv"
	"errors"
	"os"
)

func ParseCSV(filePath string) ([]Table, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return nil, errors.New("空のCSVファイルです")
	}

	tables := make(map[string]*Table)
	for i, record := range records {
		if i == 0 {
			continue
		}

		if len(record) < 10 {
			return nil, errors.New("不正なCSVファイルです。カラムが足りません。")
		}

		tableName := record[0]
		table, exists := tables[tableName]
		if !exists {
			table = &Table{Name: tableName}
			tables[tableName] = table
		}

		column := Column{
			Name:             record[1],
			Type:             record[2],
			IsPrimaryKey:     ParseBool(record[3]),
			IsUnique:         ParseBool(record[4]),
			IsNotNull:        ParseBool(record[5]),
			ForeignKeyTable:  record[6],
			ForeignKeyColumn: record[7],
			Check:            record[8],
			Comment:          record[9],
		}

		table.Columns = append(table.Columns, column)
	}

	tableSlice := make([]Table, 0, len(tables))
	for _, table := range tables {
		tableSlice = append(tableSlice, *table)
	}

	return tableSlice, nil
}
