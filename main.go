// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	src := readInput()
	if src == "" {
		fmt.Println(0)
	} else {
		words := strings.Split(src, " ")
		fmt.Println(len(words))
	}
}

// readInput reads source string
// from command line arguments and returns them.
func readInput() (src string) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	return src
}
