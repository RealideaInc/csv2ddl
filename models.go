package main

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name             string
	Type             string
	IsPrimaryKey     bool
	IsUnique         bool
	IsNotNull        bool
	Comment          string
	Check            string
	ForeignKeyTable  string
	ForeignKeyColumn string
}
