package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	var s scanner.Scanner
	input := "x := 5 + 3\nif x > 5 {\n    fmt.Println(\"x is greater than 5\")\n}"
	s.Init(strings.NewReader(input))
	for {
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}
		fmt.Printf("Token: %s, Rune: %c\n", s.TokenText(), tok)
	}
}