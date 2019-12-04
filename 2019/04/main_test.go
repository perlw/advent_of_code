package main

import (
	"testing"
)

func TestPasswordCheck1(t *testing.T) {
	tests := []struct {
		p int
		e bool
	}{
		{
			p: 111111,
			e: true,
		},
		{
			p: 223450,
			e: false,
		},
		{
			p: 123789,
			e: false,
		},
	}

	for _, test := range tests {
		got := CheckPassword1(test.p)
		if got != test.e {
			t.Errorf("Password %d, expected %v, got %v", test.p, test.e, got)
		}
	}
}

func TestPasswordCheck2(t *testing.T) {
	tests := []struct {
		p int
		e bool
	}{
		{
			p: 112233,
			e: true,
		},
		{
			p: 123444,
			e: false,
		},
		{
			p: 111122,
			e: true,
		},
		{
			p: 114444,
			e: true,
		},
	}

	for _, test := range tests {
		got := CheckPassword2(test.p)
		if got != test.e {
			t.Errorf("Password %d, expected %v, got %v", test.p, test.e, got)
		}
	}
}
