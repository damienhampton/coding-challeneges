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

func TestThreeWords(t *testing.T) {
	result := Wordcount("text text text")
	if result != "1 3 14" {
		t.Errorf("Wordcount(\"text\") = %s; want 1 3 14", result)
	}
}

func TestTwoLines(t *testing.T) {
	result := Wordcount("text\ntext")
	if result != "2 2 9" {
		t.Errorf("Wordcount(\"text\") = %s; want 2 2 8", result)
	}
}
