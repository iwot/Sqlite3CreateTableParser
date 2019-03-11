package parser_test

import (
	"testing"

	"github.com/iwot/Sqlite3CreateTableParser/parser"
)

func TestParseCase1(t *testing.T) {
	query := `CREATE TABLE members (
		id PRIMARY KEY,
		name text NOT NULL,
		type integer NOT NULL DEFAULT 0
	  )`
	table, errCode := parser.ParseTable(query, 0)
	if errCode != parser.ERROR_NONE {
		t.Fatalf("failed %#v", errCode)
	}

	found := false
	for _, column := range table.Columns {
		if column.Name == "type" {
			found = true
			if column.DefaultExpr != "0" {
				t.Fatalf("column.DefaultExpr miss match %#v", column.DefaultExpr)
			}
		}
	}

	if !found {
		t.Fatalf("column.DefaultExpr not found")
	}
}

func TestParseCase2(t *testing.T) {
	query := `CREATE TABLE member_items (
		id PRIMARY KEY,
		member_id integer NOT NULL,
		enable integer NOT NULL DEFAULT 0,
		item_id integer NOT NULL,
		num integer NOT NULL DEFAULT 0,
		FOREIGN KEY (member_id) REFERENCES members (id)
	  )`
	_, errCode := parser.ParseTable(query, 0)
	if errCode != parser.ERROR_NONE {
		t.Fatalf("failed %#v", errCode)
	}
}
