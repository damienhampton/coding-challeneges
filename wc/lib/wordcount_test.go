package wc

import "testing"

func TestOneWord(t *testing.T) {
	result := Wordcount("text")
	if result != "1 1 4" {
		t.Errorf("Wordcount(\"text\") = %s; want 1 1 4", result)
	}
}

func TestTwoWords(t *testing.T) {
	result := Wordcount("text text")
	if result != "1 2 9" {
		t.Errorf("Wordcount(\"text\") = %s; want 1 2 9", result)
	}
}
