package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bakardos/disgoft"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Usage: utftab begin end\n")
		os.Exit(1)
	} else {
		begin, err := strconv.ParseUint(os.Args[1], 16, 21)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		end, err := strconv.ParseUint(os.Args[2], 16, 21)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		if begin >= end {
			fmt.Print("Usage: begin < end")
			os.Exit(1)
		}
		disgoft.PrintUnicodeTable(rune(begin), rune(end))
	}
}
