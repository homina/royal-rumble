package roman

import (
	"errors"
	"regexp"
	"strings"
)

var (
	// ErrInvalidRoman occured when romanString is not valid roman numeric
	ErrInvalidRoman = errors.New("Invalid roman structure/format")
	symbols         = []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
)

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// RomanType hold the romanstring
type romanType struct {
	romanString string
}

// New return predefined roman type
func New(romanString string) romanType {
	return romanType{
		romanString: romanString,
	}
}

// Value is a func to get numeric value of romanString
func (r romanType) Value() (int, error) {
	if !r.validation() {
		return 0, ErrInvalidRoman
	}

	c := strings.Split(r.romanString, "")
	value := r.compute(c)
	return value, nil
}

func (r romanType) compute(c []string) int {
	lastDigit := 1000
	latin := 0
	for _, v := range c {
		var digit int
		digit = romanMap[v]
		if lastDigit < digit {
			latin -= 2 * lastDigit
		}
		lastDigit = digit
		latin += lastDigit
	}
	return latin
}

func (r romanType) validation() bool {
	if r.romanString == "" {
		return false
	}

	if check, _ := regexp.MatchString("^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", r.romanString); check == true {
		return check
	}
	return false
}
