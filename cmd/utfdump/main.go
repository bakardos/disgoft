package main

import (
	"fmt"
	"os"

	"github.com/bakardos/disgoft"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Print("Usage: utfdump [filename]\n")
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		disgoft.PrintBinaryDump("standard input", os.Stdin)
	} else {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Print(err, "\n")
			os.Exit(1)
		}
		defer f.Close()
		disgoft.PrintBinaryDump(os.Args[1], f)
	}
}
