package sekwence

import (
	"fmt"
)

var (
	asciiLowerCase = []rune("abcdefghijklmnopqrstuvwxyz")
	asciiUpperCase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphaNum       = []rune("0123456789")
)

func runeIndex(slice []rune, item rune) int {
	for i, elem := range slice {
		if elem == item {
			return i
		}
	}
	return -1
}

func reverseRuneSlice(slice []rune) {
	for i := 0; i < len(slice)/2; i++ {
		_t := slice[i]
		slice[i] = slice[len(slice)-1-i]
		slice[len(slice)-1-i] = _t
	}
}

func getAlphabeth(sym rune) *[]rune {
	if sym >= '0' && sym <= '9' {
		return &alphaNum

	}
	if sym >= 'A' && sym <= 'Z' {
		return &asciiUpperCase
	}
	if sym >= 'a' && sym <= 'z' {
		return &asciiLowerCase
	}
	return nil
}

func symbolSucc(sym rune) (next rune, carry bool, alphabeth *[]rune, err error) {
	err = nil
	alphabeth = getAlphabeth(sym)
	if alphabeth == nil {
		return 0, false, nil, fmt.Errorf("no suitable alphabeth found for %v", sym)
	}

	// get index of the next symbol in the alphabeth
	i := runeIndex(*alphabeth, sym) + 1

	if i == len(*alphabeth) {
		// if alphabeth is over, set the carry flag and start from the beginning
		carry = true
		i = 0
	} else {
		// otherwise there's nothing to carry
		carry = false
	}
	next = (*alphabeth)[i]
	return
}

// Succ implements ruby-like String.succ
func Succ(s string) (string, error) {

	var (
		symbolList    []rune
		i             int
		sym           rune
		next          rune
		err           error
		currAlphabeth *[]rune
		carry         bool
	)

	if len(s) == 0 {
		return "", nil
	}

	symbolList = []rune(s)
	reverseRuneSlice(symbolList)

	i = 0
	for {
		sym = symbolList[i]
		next, carry, currAlphabeth, err = symbolSucc(sym)
		if err != nil {
			return "", err
		}

		symbolList[i] = next
		if !carry {
			break
		}
		i++

		if i == len(symbolList) {
			if currAlphabeth == &alphaNum {
				// if carrying numbers, the next position value is usually rather 1 than 0
				symbolList = append(symbolList, (*currAlphabeth)[1])
			} else {
				// otherwise just take the first symbol of the alphabeth and don't sweat it
				symbolList = append(symbolList, (*currAlphabeth)[0])
			}
			break
		}
	}

	reverseRuneSlice(symbolList)
	return string(symbolList), nil
}
