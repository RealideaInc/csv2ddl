package main

import (
	"fmt"
	"strings"
)

func GenerateDDL(tables []Table) (string, error) {
	ddl := []string{}

	for _, table := range tables {
		tableDDL, err := generateTableDDL(&table)
		if err != nil {
			return "", err
		}
		ddl = append(ddl, tableDDL)
	}

	return strings.Join(ddl, "\n\n"), nil
}

func generateTableDDL(table *Table) (string, error) {
	if len(table.Columns) == 0 {
		return "", fmt.Errorf("テーブル %s にカラムがありません。", table.Name)
	}

	columnDDLs := []string{}
	primaryKeys := []string{}
	for _, column := range table.Columns {
		columnDDL, err := generateColumnDDL(&column)
		if err != nil {
			return "", err
		}
		columnDDLs = append(columnDDLs, columnDDL)

		if column.IsPrimaryKey {
			primaryKeys = append(primaryKeys, column.Name)
		}
	}

	ddl := fmt.Sprintf("CREATE TABLE %s (\n", table.Name)
	ddl += strings.Join(columnDDLs, ",\n")

	if len(primaryKeys) > 0 {
		ddl += fmt.Sprintf(",\nPRIMARY KEY (%s)", strings.Join(primaryKeys, ", "))
	}

	ddl += "\n);"

	return ddl, nil
}

func generateColumnDDL(column *Column) (string, error) {
	ddl := fmt.Sprintf("  %s %s", column.Name, column.Type)

	if column.IsNotNull {
		ddl += " NOT NULL"
	}

	if column.IsUnique {
		ddl += " UNIQUE"
	}

	if column.Check != "" {
		ddl += fmt.Sprintf(" CHECK (%s)", column.Check)
	}

	if column.ForeignKeyTable != "" && column.ForeignKeyColumn != "" {
		ddl += fmt.Sprintf(" FOREIGN KEY REFERENCES %s(%s)", column.ForeignKeyTable, column.ForeignKeyColumn)
	}

	if column.Comment != "" {
		ddl += fmt.Sprintf(" COMMENT '%s'", column.Comment)
	}

	return ddl, nil
}
