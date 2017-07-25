package testcalc_test

import (
	"testing"

	"github.com/worzeel/testcalc"
)

func TestIShouldBeAbleToCallAddToObtainKnownValues(t *testing.T) {

	for _, theTests := range []struct {
		numbers       []int
		expectedValue int
	}{
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3}, 6},
		{[]int{1, 5, 2}, 8},
	} {
		val := testcalc.Add(theTests.numbers...)

		if val != theTests.expectedValue {
			t.Errorf("Value should be %d but was %d", theTests.expectedValue, val)
		}
	}
}

func TestIShouldBeAbleToCallSubtractPassingInMultipleValues(t *testing.T) {
	testcalc.Subtract(1, 2)
}
