// challenges/testing/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ &/)", want: 1},
		"two":   {input: "812 %6ab//", want: 2},
	}

	lc := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := lc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count

func TestNumberCount(t *testing.T) {
	var tests = map[string]struct {
		input    string
		want     int
		wantname string
	}{
		"empty": {input: "", want: 0, wantname: "testname"},
		"one":   {input: "abc 1,?/", want: 1, wantname: "testname"},
		"two":   {input: "abc 12&8 ^", want: 3, wantname: "testname"},
	}

	nc := numberCounter{"testname"}
	// nametest
	fmt.Println("***designation***")
	fmt.Println(nc.name())

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := nc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
			if gotname := nc.name(); gotname != tc.wantname {
				t.Errorf("gotname %v wantname %v", gotname, tc.wantname)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "abc 1,?/", want: 4},
		"two":   {input: "abc 12&8 ^_", want: 5},
	}

	sc := symbolCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := sc.count(tc.input); got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("***SETUP***")

	m.Run()

	fmt.Println("***TEARDOWN***")
}
