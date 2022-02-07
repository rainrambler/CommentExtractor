package main

import (
	"testing"
)

func TestFormatString1(t *testing.T) {
	s := "// aa\nint fn(){}\n"

	res := formatString(s)
	expected := " aa\n"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}

func TestFormatString2(t *testing.T) {
	s := "/* *aa**/\nint fn(){}\n"

	res := formatString(s)
	expected := " *aa*"

	if res != expected {
		t.Errorf("Result: %v, want: %v", res, expected)
	}
}
