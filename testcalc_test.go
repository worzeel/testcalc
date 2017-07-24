package testcalc_test

import (
	"testing"

	"github.com/worzeel/testcalc"
)

func TestIShouldBeAbleToCallAddPassingInDifferentAmountOfValues(t *testing.T) {
	testcalc.Add(1, 2)
	testcalc.Add(1, 2, 3)
}

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
