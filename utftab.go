package main

import (
	"fmt"
	"os"
	"strconv"
)

// PrintUnicodeTable печатает таблицу.
func PrintUnicodeTable(begin, end rune) {
	for num := begin; num < end; num++ {
		symbol := string(num)
		dump := "BINARY DUMP"
		fmt.Printf("%5X  %s  %s\n", num, dump, symbol)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Usage: utftab begin end")
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
		PrintUnicodeTable(rune(begin), rune(end))
	}
}
