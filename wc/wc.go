package main

import (
	wc "damien/wc/lib"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	contents, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Couldn't read file %s\n", err)
	}

	fmt.Printf("%s %s\n", wc.Wordcount(string(contents)), filename)
}
