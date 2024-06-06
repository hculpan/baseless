package dbf_test

import (
	"baseless/pkg/dbf"
	"testing"
)

func TestNewDbf(t *testing.T) {
	d, err := dbf.NewDbf("")
	if err != nil {
		t.Fatal(err)
	} else if d == nil {
		t.Fatal("expected Dbf object, got nil")
	}
}
