package stringCalculator

import (
	"strconv"
	"strings"
)

func StringCalculator(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}

	//split := strings.Split(s, ",")
	split := strings.FieldsFunc(s, Split)
	sum := 0
	for _, val := range split {
		atoi, err := strconv.Atoi(val)
		if err != nil {
			return -1, err
		}
		sum += atoi
	}

	return sum, nil
}

func Split(r rune) bool {
	return r == ',' || r == '\n'
}
