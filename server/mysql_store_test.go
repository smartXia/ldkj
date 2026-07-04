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

func TestBannerOrderClauseFallsBackForLegacySchema(t *testing.T) {
	if got := bannerOrderClause(false); got != "ORDER BY id ASC" {
		t.Fatalf("expected legacy banner order fallback, got %q", got)
	}
	if got := bannerOrderClause(true); got != "ORDER BY sort_order ASC, id ASC" {
		t.Fatalf("expected sort_order banner order, got %q", got)
	}
}
