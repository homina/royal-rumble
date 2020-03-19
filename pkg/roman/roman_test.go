package roman

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRoman(t *testing.T) {
	str := "hello"
	var roman interface{}
	roman = New(str)
	data, ok := roman.(romanType)
	assert.Equal(t, ok, true, "new func should return roman type")
	assert.Equal(t, data.romanString, str, `new func should have "`+str+`" value when called with "`+str+`"`)
}

func TestValidRoman(t *testing.T) {
	testLists := map[string]bool{
		"":      false,
		"XXX69": false,
		"ML":    true,
		"ml":    false,
	}
	// you can continue to make test list if you want

	for k, v := range testLists {
		roman := New(k)
		assert.Equal(t, roman.validation(), v, "invalid validation")
	}

}

func TestCompute(t *testing.T) {
	testLists := map[string]int{
		"XXV":        25,
		"LXXVII":     77,
		"CXXXVIII":   138,
		"CDXVIII":    418,
		"CMLXXXVIII": 988,
	}

	for k, v := range testLists {
		roman := New(k)
		value := roman.compute(strings.Split(k, ""))
		assert.Equal(t, value, v, "invalid compute")
	}

}

func TestValue(t *testing.T) {
	testListsInvalid := map[string]error{
		"":               ErrInvalidRoman,
		"XXX69":          ErrInvalidRoman,
		"glob glob fyuh": ErrInvalidRoman,
		"globglobfyuh":   ErrInvalidRoman,
		"XXV":            nil,
	}

	for k, v := range testListsInvalid {
		roman := New(k)
		_, err := roman.Value()
		assert.Equal(t, err, v, "invalid compute error message")
	}

	testLists := map[string]int{
		"XXV":        25,
		"LXXVII":     77,
		"CXXXVIII":   138,
		"CDXVIII":    418,
		"CMLXXXVIII": 988,
	}

	for k, v := range testLists {
		roman := New(k)
		value, _ := roman.Value()
		assert.Equal(t, value, v, "invalid compute result")
	}

}
