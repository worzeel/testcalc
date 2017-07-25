package testcalc_test

import (
	"testing"

	"github.com/worzeel/testcalc"
)

type calcChecks func(...int) int
type values []struct {
	numbers       []int
	expectedValue int
}

func doMultipleChecks(fn calcChecks, vals values, t *testing.T) {
	for _, theTests := range vals {
		val := fn(theTests.numbers...)

		if val != theTests.expectedValue {
			t.Errorf("Value should be %d but was %d", theTests.expectedValue, val)
		}
	}
}

func TestIShouldBeAbleToCallAddToObtainKnownValues(t *testing.T) {
	doMultipleChecks(testcalc.Add, values{
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3}, 6},
		{[]int{1, 5, 2}, 8},
	}, t)
}

func TestIShouldBeAbleToCallSubtractToObtainKnownValues(t *testing.T) {
	doMultipleChecks(testcalc.Subtract, values{
		{[]int{1, 2}, -1},
		{[]int{3, 2, 1}, 0},
		{[]int{10, 5, 2}, 3},
	}, t)
}

func TestIShouldBeAbleToCallMultiplyToObtainKnownValues(t *testing.T) {
	doMultipleChecks(testcalc.Multiply, values{
		{[]int{1, 2}, 2},
		{[]int{3, 2, 1}, 6},
		{[]int{10, 5, 2}, 100},
	}, t)
}
