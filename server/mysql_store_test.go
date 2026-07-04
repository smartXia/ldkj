package main

import (
	"strings"
	"testing"
)

func TestServiceSelectColumnsCoalesceNullableTextFields(t *testing.T) {
	nullableColumns := []string{
		"subtitle",
		"summary",
		"cover_url",
		"icon_url",
		"content",
		"highlights",
		"process_steps",
	}

	for _, column := range nullableColumns {
		expected := "COALESCE(" + column + ", '')"
		if !strings.Contains(serviceSelectColumns, expected) {
			t.Fatalf("expected serviceSelectColumns to contain %q, got %s", expected, serviceSelectColumns)
		}
	}
}
