package roman

import (
	"sort"
	"strings"
)

type byRoman []string

func (s byRoman) Len() int {
	return len(s)
}
func (s byRoman) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byRoman) Less(i, j int) bool {
	if s.extractText(s[i]) == s.extractText(s[j]) {
		return s.extractRoman(s[i]) < s.extractRoman(s[j])
	} else {
		return s.extractText(s[i]) < s.extractText(s[j])
	}
}

func (s byRoman) extractRoman(text string) int {
	fl := strings.Split(text, " ")
	f := fl[len(fl)-1]
	fv := New(f)
	first, err := fv.Value()
	if err != nil {
		panic(err)
	}

	return first
}

func (s byRoman) extractText(text string) string {
	fl := strings.Split(text, " ")
	f := fl[:len(fl)-1]
	return strings.Join(f, " ")
}

// sort by roman
func Sort(data []string) {
	sort.Sort(byRoman(data))
}
