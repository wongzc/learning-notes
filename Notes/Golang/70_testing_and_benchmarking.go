//https://gobyexample.com/testing-and-benchmarking
package main

import (
    "fmt"
    "testing"
)
// usually testing code is in same package of the code that need to test

func IntMin(a, b int) int { 
	// the code to be tested
	// usually test code in a sourcefile like abc.go, test file will be abc_test.go
    if a < b {
        return a
    }
    return b
}

func TestIntMinBasic(t *testing.T) {
	//test function: name begin with "Test"
    ans := IntMin(2, -2)
    if ans != -2 {

        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
		// t.Error: report failure but continue test
		// t.Fatal: report test failure and stop
    }
}

func TestIntMinTableDriven(t *testing.T) {
	// use table style to avoid repetitive
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
		// t.Run: running subtest
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

func BenchmarkIntMin(b *testing.B) {
	// benchmark test, with name: Benchmark
	// to test speed/ efficiency
    for b.Loop() {

        IntMin(1, 2)
    }
}