package sekwence

import (
	"testing"
)

type stringRangeCase struct {
	from    string
	to      string
	exclude bool
	result  []string
}

var (
	succCases = map[string]string{
		"0":  "1",
		"3a": "3b",
		"9z": "10a",
		"A9": "B0",
		"ZZ": "AAA",
	}

	stringRangeCases = []stringRangeCase{
		{"a0", "b4", false, []string{"a0", "a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", "a9", "b0", "b1", "b2", "b3", "b4"}},
		{"7a", "8z", false, []string{"7a", "7b", "7c", "7d", "7e", "7f", "7g", "7h", "7i", "7j", "7k", "7l", "7m", "7n", "7o", "7p", "7q", "7r", "7s", "7t", "7u", "7v", "7w", "7x", "7y", "7z", "8a", "8b", "8c", "8d", "8e", "8f", "8g", "8h", "8i", "8j", "8k", "8l", "8m", "8n", "8o", "8p", "8q", "8r", "8s", "8t", "8u", "8v", "8w", "8x", "8y", "8z"}},
	}

	expandSinglePatternCases = map[string][]string{
		"{a0..b4}":    []string{"a0", "a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", "a9", "b0", "b1", "b2", "b3", "b4"},
		"{00..03,12}": []string{"00", "01", "02", "03", "12"},
	}
)

func TestGetAlphabeth(t *testing.T) {

	var found bool
	var alphabeth *[]rune
	var current rune

	current = 's'
	alphabeth = getAlphabeth(current)
	found = false
	for _, sym := range *alphabeth {
		if sym == current {
			found = true
			break
		}
	}
	if !found {
		t.Errorf(`symbol "%v" should be in asciiLowerCase alphabeth`, current)
	}

	current = 'Z'
	alphabeth = getAlphabeth(current)
	found = false
	for _, sym := range *alphabeth {
		if sym == current {
			found = true
			break
		}
	}
	if !found {
		t.Errorf(`symbol "%v" should be in asciiUpperCase alphabeth`, current)
	}

	current = '4'
	alphabeth = getAlphabeth(current)
	found = false
	for _, sym := range *alphabeth {
		if sym == current {
			found = true
			break
		}
	}
	if !found {
		t.Errorf(`symbol "%v" should be in alphaNum alphabeth`, current)
	}

}

func TestSymbolSucc(t *testing.T) {
	var item rune
	var carry bool
	var err error

	item, carry, _, err = symbolSucc('s')
	if item != 't' {
		t.Errorf("the succ of 's' should be 't', got %v instead", item)
	}
	if carry {
		t.Errorf("the succ of 's' should not set the carry flag")
	}
	if err != nil {
		t.Errorf("error getting symbolSucc('s'): %s", err)
	}

	item, carry, _, err = symbolSucc('Z')
	if item != 'A' {
		t.Errorf("the succ of 'Z' should be 'A', got %v instead", item)
	}
	if !carry {
		t.Errorf("the succ of 'Z' should set the carry flag")
	}
	if err != nil {
		t.Errorf("error getting symbolSucc('Z'): %s", err)
	}

	item, carry, _, err = symbolSucc('4')
	if item != '5' {
		t.Errorf("the succ of '4' should be '5', got %v instead", item)
	}
	if carry {
		t.Errorf("the succ of '4' should not set the carry flag")
	}
	if err != nil {
		t.Errorf("error getting symbolSucc('4'): %s", err)
	}

	_, _, _, err = symbolSucc('Ё')
	if err == nil {
		t.Errorf("symbolSucc('Ё') should throw an error")
	}

}

func TestReverseSlice(t *testing.T) {
	s := []rune{'a', 'b', 'c', 'd', 'e'}
	rs := []rune{'e', 'd', 'c', 'b', 'a'}

	reverseRuneSlice(s)
	for i := 0; i < len(s); i++ {
		if s[i] != rs[i] {
			t.Errorf("Invalid symbol at pos %d after reverse: expected %v but got %v", i, rs[i], s[i])
		}
	}
}

func TestSucc(t *testing.T) {
	for arg, exp := range succCases {
		res, err := Succ(arg)
		if err != nil {
			t.Errorf(`Error during Succ("%s"): %s`, arg, err)
		}
		if res != exp {
			t.Errorf(`Succ("%s") is expected to be "%s" but got "%s"`, arg, exp, res)
		}
	}
}

func TestStringRange(t *testing.T) {
	for _, tc := range stringRangeCases {
		res, err := StringRange(tc.from, tc.to, tc.exclude)
		if err != nil {
			t.Errorf("Error during StringRange(%v, %v, %v): %s", tc.from, tc.to, tc.exclude, err)
		}
		for i := 0; i < len(res); i++ {
			if res[i] != tc.result[i] {
				t.Errorf("Invalid StringRange item at pos %d: expected %v but got %v", i, tc.result[i], res[i])
			}
		}
	}
}

func TestExpandSinglePattern(t *testing.T) {
	for arg, exp := range expandSinglePatternCases {
		res, err := expandSinglePattern(arg)
		if err != nil {
			t.Errorf("Error during ExpandSinglePattern(%v): %s", arg, err)
		}
		for i := 0; i < len(res); i++ {
			if res[i] != exp[i] {
				t.Errorf("Invalid ExpandSinglePattern result item at pos %d: expected %v but got %v", i, exp[i], res[i])
			}
		}
	}
}
