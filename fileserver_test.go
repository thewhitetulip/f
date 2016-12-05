package main

import (
	"testing"
)

func TestPortNumber(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{0, false},
		{1, true},
		{65535, true},
		{65536, false},
		{-100, false},
	}

	for _, test := range tests {
		if got := IsValidPort(test.input); got != test.want {
			t.Errorf("IsValidPortnumber(%v) = %v", test.input, got)
		}
	}
}
