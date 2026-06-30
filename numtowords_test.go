package numtowords_test

import (
	"testing"

	numtowords "github.com/vishwa-100/numtowords"
)

func TestInvalidNumber(t *testing.T) {
	_, err := numtowords.ConvertToWords(numtowords.MaxNum + 1)
	if err == nil {
		t.Log("Expected error for number greater than MaxNum, but got nil")
		t.Fail()
	}

	_, err = numtowords.ConvertToWords(numtowords.MinNum - 1)
	if err == nil {
		t.Log("Expected error for number less than MinNum, but got nil")
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	result, err := numtowords.ConvertToWords(0)
	if err != nil {
		t.Logf("Convert with zero returned error: %v", err)
		t.FailNow()
	}

	if result != "zero" {
		t.Logf("Convert with zero returned %q, expected %q", result, "zero")
		t.Fail()
	}
}

func TestUnits(t *testing.T) {

	units := [20]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	for index, value := range units {
		t.Logf("Now testing %v", index)
		result, err := numtowords.ConvertToWords(index)
		if err != nil {
			t.Logf("Convert with %d returned error: %v", index, err)
			t.Fail() // Not FailNow as we have multiple test cases to run
		}

		if result != value {
			t.Logf("Convert with %d returned %q, expected %q", index, result, value)
			t.Fail()
		}
	}
}

func TestTens(t *testing.T) {
	tens := [8]string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	for index, value := range tens {
		num := (index + 2) * 10
		t.Logf("Now testing %v", num)
		result, err := numtowords.ConvertToWords(num)
		if err != nil {
			t.Logf("Convert with %d returned error: %v", num, err)
			t.Fail() // Not FailNow as we have multiple test cases to run
		}

		if result != value {
			t.Logf("Convert with %d returned %q, expected %q", num, result, value)
			t.Fail()
		}
	}
}

func TestHundreds(t *testing.T) {
	testcases := map[int]string{
		100: "one hundred",
		201: "two hundred and one",
		420: "four hundred and twenty",
		333: "three hundred and thirty three",
	}

	for num, value := range testcases {
		t.Logf("Now testing %v", num)
		result, err := numtowords.ConvertToWords(num)
		if err != nil {
			t.Logf("Convert with %d returned error: %v", num, err)
			t.Fail() // Not FailNow as we have multiple test cases to run
		}

		if result != value {
			t.Logf("Convert with %d returned %q, expected %q", num, result, value)
			t.Fail()
		}
	}
}

func TestNegativeNumbers(t *testing.T) {
	testcases := map[int]string{
		-1:   "minus one",
		-50:  "minus fifty",
		-18:  "minus eighteen",
		-102: "minus one hundred and two",
		-218: "minus two hundred and eighteen",
		-420: "minus four hundred and twenty",
		-333: "minus three hundred and thirty three",
	}

	for num, value := range testcases {
		t.Logf("Now testing %v", num)
		result, err := numtowords.ConvertToWords(num)
		if err != nil {
			t.Logf("Convert with %d returned error: %v", num, err)
			t.Fail() // Not FailNow as we have multiple test cases to run
		}

		if result != value {
			t.Logf("Convert with %d returned %q, expected %q", num, result, value)
			t.Fail()
		}
	}
}
