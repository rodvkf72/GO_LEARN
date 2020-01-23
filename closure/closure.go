package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string

	err := ReadFrom(r, func(line string) {
		lines = append(lines, line)
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}
}

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
