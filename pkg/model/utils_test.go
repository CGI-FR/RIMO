package model_test

import (
	"testing"

	"github.com/cgi-fr/rimo/pkg/model"
)

func TestAddColumn(t *testing.T) {
	t.Parallel()

	base := model.NewBase("test_base")

	column := model.Column{ //nolint:exhaustruct
		Name:    "test_column",
		Type:    model.ColType.String,
		Concept: "test_concept",
	}

	tableName := "test_table"

	base.AddColumn(column, tableName)

	// fmt.Print(valast.String(base))

	if len(base.Tables) != 1 {
		t.Errorf("expected 1 table, got %d", len(base.Tables))
	}

	if base.Tables[0].Name != tableName {
		t.Errorf("expected table name %q, got %q", tableName, base.Tables[0].Name)
	}

	if len(base.Tables[0].Columns) != 1 {
		t.Errorf("expected 1 column, got %d", len(base.Tables[0].Columns))
	}

	if base.Tables[0].Columns[0].Name != column.Name {
		t.Errorf("expected column name %q, got %q", column.Name, base.Tables[0].Columns[0].Name)
	}

	if base.Tables[0].Columns[0].Type != column.Type {
		t.Errorf("expected column type %q, got %q", column.Type, base.Tables[0].Columns[0].Type)
	}

	if base.Tables[0].Columns[0].Concept != column.Concept {
		t.Errorf("expected column concept %q, got %q", column.Concept, base.Tables[0].Columns[0].Concept)
	}
}
