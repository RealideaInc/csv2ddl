package main

import (
	"fmt"
	"strings"
)

func GenerateDDL(tables []Table) (string, error) {
	var ddl strings.Builder

	for _, table := range tables {
		ddl.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", table.Name))

		var columnDeclarations []string
		var primaryKeyColumns []string
		for _, column := range table.Columns {
			columnDeclaration := fmt.Sprintf("  %s %s", column.Name, column.Type)

			if column.IsPrimaryKey {
				primaryKeyColumns = append(primaryKeyColumns, column.Name)
			}

			if column.IsNotNull {
				columnDeclaration += " NOT NULL"
			}

			if column.IsUnique {
				columnDeclaration += " UNIQUE"
			}

			if column.Check != "" {
				columnDeclaration += fmt.Sprintf(" CHECK(%s)", column.Check)
			}

			if column.Comment != "" {
				columnDeclaration += fmt.Sprintf(" COMMENT '%s'", column.Comment)
			}

			columnDeclarations = append(columnDeclarations, columnDeclaration)
		}

		if len(primaryKeyColumns) > 0 {
			primaryKeyDeclaration := fmt.Sprintf("PRIMARY KEY (%s)", strings.Join(primaryKeyColumns, ", "))
			columnDeclarations = append(columnDeclarations, primaryKeyDeclaration)
		}

		for _, column := range table.Columns {
			if column.ForeignKeyTable != "" && column.ForeignKeyColumn != "" {
				foreignKeyDeclaration := fmt.Sprintf("FOREIGN KEY (%s) REFERENCES %s(%s)", column.Name, column.ForeignKeyTable, column.ForeignKeyColumn)
				columnDeclarations = append(columnDeclarations, foreignKeyDeclaration)
			}
		}

		ddl.WriteString(strings.Join(columnDeclarations, ",\n"))
		ddl.WriteString("\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;\n")
	}

	if ddl.Len() == 0 {
		return "", fmt.Errorf("生成されたDDLが空です")
	}

	return ddl.String(), nil
}
