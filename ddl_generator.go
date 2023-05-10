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
	foreignKeys := []string{}
	for _, column := range table.Columns {
		columnDDL, err := generateColumnDDL(&column)
		if err != nil {
			return "", err
		}
		columnDDLs = append(columnDDLs, columnDDL)

		if column.IsPrimaryKey {
			primaryKeys = append(primaryKeys, column.Name)
		}
		if column.ForeignKeyTable != "" && column.ForeignKeyColumn != "" {
			fkr := fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s(%s)", column.Name, column.ForeignKeyTable, column.ForeignKeyColumn)
			foreignKeys = append(foreignKeys, fkr)
		}
	}

	ddl := fmt.Sprintf("CREATE TABLE %s (\n", table.Name)
	ddl += strings.Join(columnDDLs, ",\n")

	if len(primaryKeys) > 0 {
		ddl += fmt.Sprintf(",\nPRIMARY KEY (%s)", strings.Join(primaryKeys, ", "))
	}

	if len(foreignKeys) > 0 {
		ddl += ",\n"
		ddl += strings.Join(foreignKeys, ",\n")
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

	if column.Comment != "" {
		ddl += fmt.Sprintf(" COMMENT '%s'", column.Comment)
	}

	return ddl, nil
}
