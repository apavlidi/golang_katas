package stringCalculator

import (
	"strconv"
	"strings"
)

const DefaultDelimiters = ",\n"

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
