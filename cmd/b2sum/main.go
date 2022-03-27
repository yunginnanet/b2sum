package main

import (
	"fmt"
	"io"
	"os"

	"clake2/internal"
)

var input io.Reader

func init() {
	if len(os.Args) < 2 {
		input = os.Stdin
	} else {
		var err error
		input, err = os.Open(os.Args[1])
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
}

func main() {
	_, err := fmt.Fprintf(os.Stdout, internal.C264(input))
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
