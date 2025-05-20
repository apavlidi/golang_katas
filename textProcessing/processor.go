package textProcessing

import (
	"slices"
	"sort"
	"strings"
)

const (
	Delimiter      = " "
	MaxCommonLimit = 10
)

type Processor interface {
	analyse(text string) []string
}

type Analytics struct {
	wordCount        int
	mostPopularWords []string
	timeToRead       ReadTime
}

type ReadTime struct {
	minute  int
	seconds int
}

type TextProcessor struct{}

func (t TextProcessor) analyse(text string) Analytics {
	wordCount := make(map[string]int)

	split := strings.Split(strings.ToLower(text), Delimiter)
	for _, val := range split {
		wordCount[val]++
	}

	var keys []string
	for k := range wordCount {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return wordCount[keys[i]] < wordCount[keys[j]]
	})
	slices.Reverse(keys)

	mostCommon := getMostCommon(keys, MaxCommonLimit)

	return Analytics{
		wordCount:        len(split),
		mostPopularWords: mostCommon,
		timeToRead:       ReadTime{},
	}
}

func getMostCommon(keys []string, limit int) []string {
	var mostCommon []string
	for _, val := range keys {
		if limit > 0 {
			mostCommon = append(mostCommon, val)
		}
		limit--
	}
	return mostCommon
}
