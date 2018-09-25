package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode"
)

func fileToString(file *os.File) string {
	buffer := bytes.Buffer{}

	file.Seek(0, 0)
	input := bufio.NewScanner(file)

	for input.Scan() {
		for _, s := range input.Text() {
			if !unicode.IsSpace(rune(s)) {
				if s != '#' {
					buffer.Write([]byte(fmt.Sprintf("%s\n", input.Text())))
				}
				break
			}
		}
	}

	return buffer.String()
}

func outputFile(file *os.File) {
	os.Stdout.Write([]byte(fileToString(file)))
}
