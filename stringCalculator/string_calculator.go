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
	sum := 0
	for _, val := range split {
		num, err := strconv.Atoi(val)
		if err != nil {
			return -1, err
		}
		sum += num
	}

	return sum, nil
}

func getCustomSeparator(s string) string {
	split := strings.Split(s, "")
	if len(split) >= 2 {
		e := []string{split[0], split[1]}
		firstTwoCharacters := strings.Join(e, "")
		if strings.Contains(firstTwoCharacters, "//") {
			return split[2]
		}
	}
	return ""
}

func makeSplitter(delimiters string) func(rune) bool {
	return func(r rune) bool {
		return strings.ContainsRune(delimiters, r)
	}
}

func defaultSplitter(r rune) func(rune) bool {
	return makeSplitter(",\n")
}
