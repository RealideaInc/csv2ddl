package main

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name             string
	Type             string
	IsPrimaryKey     bool
	IsNotNull        bool
	IsUnique         bool
	ForeignKeyTable  string
	ForeignKeyColumn string
	Check            string
	Comment          string
}
