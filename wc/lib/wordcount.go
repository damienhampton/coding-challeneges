package wc

import (
	"fmt"
	"strings"
)

func Wordcount(s string) (res string) {
	lines := len(strings.Split(s, "\n"))
	words := len(strings.Split(s, " "))
	chars := len(s)
	return fmt.Sprintf("%d %d %d", lines, words, chars)
}
