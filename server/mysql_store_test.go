package main

import (
	"strings"
	"testing"
)

func TestServiceSelectColumnsCoalesceNullableTextFields(t *testing.T) {
	columns := serviceSelectColumns(true)
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
		if !strings.Contains(columns, expected) {
			t.Fatalf("expected serviceSelectColumns to contain %q, got %s", expected, columns)
		}
	}
}

func TestServiceSelectColumnsFallsBackWhenProcessStepsMissing(t *testing.T) {
	columns := serviceSelectColumns(false)
	if strings.Contains(columns, "process_steps") {
		t.Fatalf("expected legacy service columns to avoid process_steps, got %s", columns)
	}
	if !strings.Contains(columns, "''") {
		t.Fatalf("expected legacy service columns to include empty process placeholder, got %s", columns)
	}
}

func TestBannerOrderClauseFallsBackForLegacySchema(t *testing.T) {
	if got := bannerOrderClause(false); got != "ORDER BY id ASC" {
		t.Fatalf("expected legacy banner order fallback, got %q", got)
	}
	if got := bannerOrderClause(true); got != "ORDER BY sort_order ASC, id ASC" {
		t.Fatalf("expected sort_order banner order, got %q", got)
	}
}
