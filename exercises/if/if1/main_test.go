// if1
// Make me compile!

package main_test

import (
	"math"
	"testing"
)

func bigger(a int, b int) int {
	// Complete this function to return the bigger number
	// Use only if statements
	return int(math.Max(float64(a), float64(b)))
}

func TestTwoIsBiggerThanOne(t *testing.T) {
	if bigger(2, 1) != 2 {
		t.Errorf("2 is bigger than 1")
	}
}

func TestTenIsBiggerThanFive(t *testing.T) {
	if bigger(5, 10) != 10 {
		t.Errorf("10 is bigger than 5")
	}
}
