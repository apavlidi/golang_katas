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
		s = strings.SplitAfter(s, "//")[1]
	}
	separator := makeSplitter(DefaultDelimiters + customSeparator)

	split := strings.FieldsFunc(s, separator)
	return sumInts(split)
}

func sumInts(s []string) (int, error) {
	sum := 0
	for _, val := range s {
		num, err := strconv.Atoi(val)
		if num < 0 {
			return -1, &NotAllowedNegativeNumbers{
				Input: num,
				Msg:   "Negative values are not allowed",
			}
		}
		if err != nil {
			return -1, err
		}
		sum += num
	}

	return sum, nil
}

func getCustomSeparator(s string) string {
	if strings.HasPrefix(s, "//") && len(s) >= 2 {
		return string(s[2])
	}
	return ""
}

func makeSplitter(delimiters string) func(rune) bool {
	return func(r rune) bool {
		return strings.ContainsRune(delimiters, r)
	}
}

func (e *NotAllowedNegativeNumbers) Error() string {
	return fmt.Sprintf("parse error: %s (%d)", e.Msg, e.Input)
}
