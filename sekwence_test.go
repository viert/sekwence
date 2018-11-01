package sekwence

import (
	"testing"
)

var (
	succCases = map[string]string{
		"0":  "1",
		"3a": "3b",
		"9z": "10a",
		"A9": "B0",
		"ZZ": "AAA",
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
