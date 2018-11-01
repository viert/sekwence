package sekwence

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"regexp"
	"strings"
)

type expressionIndex struct {
	start int
	end   int
}

var (
	asciiLowerCase = []rune("abcdefghijklmnopqrstuvwxyz")
	asciiUpperCase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alphaNum       = []rune("0123456789")

	singleExpr = regexp.MustCompile(`^([0-9a-zA-Z]+)\.\.([0-9a-zA-Z]+)$`)
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

// StringRange implements ruby-like string range generators
// i.e. ruby's ("a0".."e4").to_a equals to StringRange("a0", "e4", false)
// exclude params indicates that value of "to" should be excluded
//
func StringRange(from string, to string, exclude bool) ([]string, error) {
	var err error

	result := make([]string, 0)
	for from != to && len(from) <= len(to) {
		result = append(result, from)
		from, err = Succ(from)
		if err != nil {
			return result, err
		}
	}

	if from == to && !exclude {
		result = append(result, from)
	}
	return result, nil
}

func expandSinglePattern(pattern string) ([]string, error) {
	if strings.HasPrefix(pattern, "{") && strings.HasSuffix(pattern, "}") {
		pattern = pattern[1 : len(pattern)-1]
	}
	tokens := strings.Split(pattern, ",")
	result := make([]string, 0)

	for _, token := range tokens {
		corners := singleExpr.FindStringSubmatch(token)
		if len(corners) == 0 {
			result = append(result, token)
		} else {
			from := corners[1]
			to := corners[2]
			rng, err := StringRange(from, to, false)
			if err != nil {
				return nil, fmt.Errorf("can't parse expression %s", token)
			}
			result = append(result, rng...)
		}
	}
	return result, nil
}

func getBracesIndices(pattern string) ([]expressionIndex, error) {
	stack := stack.New()
	indices := make([]expressionIndex, 0)
	currentIndex := expressionIndex{start: -1, end: -1}

	for ind, sym := range pattern {
		if sym == '{' {
			stack.Push(true)
			currentIndex.start = ind
		} else if sym == '}' {
			val := stack.Pop()
			if val == nil {
				return nil, fmt.Errorf("closing brace without opening one at %d", ind)
			}
			currentIndex.end = ind
		}

		if currentIndex.end > 0 {
			if stack.Len() != 0 {
				return nil, fmt.Errorf("nested patterns are not allowed")
			}
			indices = append(indices, currentIndex)
			currentIndex = expressionIndex{start: -1, end: -1}
		}
	}
	if stack.Len() > 0 {
		return nil, fmt.Errorf("closing brace expected at the end of pattern")
	}
	return indices, nil
}

// ExpandPattern expands the whole pattern recursively
// and returns all the resulting matches
func ExpandPattern(pattern string) ([]string, error) {
	indices, err := getBracesIndices(pattern)
	if err != nil {
		return nil, err
	}

	if len(indices) == 0 {
		// no range expressions found
		return []string{pattern}, nil
	}

	results := make([]string, 0)
	ind := indices[0]
	tokens, err := expandSinglePattern(pattern[ind.start+1 : ind.end])
	if err != nil {
		return nil, err
	}

	for _, token := range tokens {
		replaced := pattern[:ind.start] + token + pattern[ind.end+1:]
		preliminary, err := ExpandPattern(replaced)
		if err != nil {
			return nil, err
		}
		results = append(results, preliminary...)
	}
	return results, nil
}
