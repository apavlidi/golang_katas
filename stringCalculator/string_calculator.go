package stringCalculator

import (
	"fmt"
	"strconv"
	"strings"
)

const DefaultDelimiters = ",\n"

type NotAllowedNegativeNumbers struct {
	Input int
	Msg   string
}

func StringCalculator(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}

	customSeparator := getCustomSeparator(s)
	if customSeparator != "" {
		if len(customSeparator) > 1 {
			s = strings.SplitAfter(s, "//["+customSeparator+"]")[1]
		} else {
			s = strings.SplitAfter(s, "//")[1]
		}
	}
	separator := makeSplitter(DefaultDelimiters + customSeparator)

	split := strings.FieldsFunc(s, separator)
	return sumInts(split)
}

func sumInts(s []string) (int, error) {
	sum := 0
	for _, val := range s {
		number, err := parseNumber(val)
		if err != nil {
			return -1, err
		}
		sum += number
	}

	return sum, nil
}

func parseNumber(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}

	if num < 0 {
		return -1, &NotAllowedNegativeNumbers{
			Input: num,
			Msg:   "Negative values are not allowed",
		}
	}

	if isBigNumber(num) {
		return 0, nil
	}
	return num, nil
}

func isBigNumber(num int) bool {
	return num > 1000
}

func getCustomSeparator(s string) string {
	if strings.HasPrefix(s, "//") && len(s) >= 2 {
		separator, hasArbitraryLengthSeparator := getArbitraryLengthSeparator(s)
		if hasArbitraryLengthSeparator {
			return separator
		}
		return string(s[2])
	}
	return ""
}

func getArbitraryLengthSeparator(s string) (string, bool) {
	separator := ""
	addToSeparator := false
	if strings.Contains(s, "[") && strings.Contains(s, "]") {
		for _, char := range s {
			if char == '[' {
				addToSeparator = true
				continue
			}

			if char == ']' {
				addToSeparator = false
				continue
			}

			if addToSeparator {
				separator += string(char)
			}
		}

		//start := strings.Index(s, "[")
		//end := strings.Index(s, "]")
		//return s[start+1 : end], true
		return separator, true
	}
	return separator, false
}

func makeSplitter(delimiters string) func(rune) bool {
	return func(r rune) bool {
		return strings.ContainsRune(delimiters, r)
	}
}

func (e *NotAllowedNegativeNumbers) Error() string {
	return fmt.Sprintf("parse error: %s (%d)", e.Msg, e.Input)
}
