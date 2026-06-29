// Package numtowords provides a function to convert numbers to their word representation.
//
// The first version converts positive integers only, from 0 to MaxNum (999).
package numtowords

import "fmt"

// MaxNum is the maximum number that can be converted to words
const MaxNum = 999

// ConvertToWords converts a number to its word representation
func ConvertToWords(num int) (string, error) {
	if num < 0 || num > MaxNum {
		return "", fmt.Errorf("can only convert numbers between 0 and %d", MaxNum)
	}
	if num == 0 {
		return "zero", nil
	}

	var result string

	units := [20]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

	tens := [8]string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if num > 99 {
		hundredsIndex := num / 100

		result += units[hundredsIndex] + " hundred"
		num = num % 100

		if num == 0 {
			return result, nil
		}

		if num > 0 {
			result += " and "
		}
	}

	if num > 19 {
		tensIndex := num/10 - 2
		result += tens[tensIndex]
		num = num % 10

		if num == 0 {
			return result, nil
		}

		result += " "
	}

	result += units[num]

	return result, nil
}
