package wc

import (
	"fmt"
	"strings"
)

func Wordcount(s string) (res string) {
	lines := strings.Split(strings.Trim(s, "\n"), "\n")
	wordCount := 0
	for i := 0; i < len(lines); i++ {
		wordCount = wordCount + len(strings.Split(lines[i], " "))

	}
	return fmt.Sprintf("%d %d %d", len(lines), wordCount, len(s))
}
