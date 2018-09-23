package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func outputFile(file *os.File) {
	file.Seek(0, 0)
	input := bufio.NewScanner(file)

	for input.Scan() {
		for _, s := range input.Text() {
			if !unicode.IsSpace(rune(s)) {
				if s != '#' {
					fmt.Printf("%s\n", input.Text())
				}
				break
			}
		}
	}

	fmt.Printf("\n")
}
