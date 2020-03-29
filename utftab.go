package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DumpRuneBinary Даёт Dump
func DumpRuneBinary(s string) string {
	var r strings.Builder
	for i := 0; i < 4; i++ {
		if i != 0 {
			r.WriteString(" ")
		}
		if i >= len(s) {
			r.WriteString(strings.Repeat(" ", 8))
		} else {
			fmt.Fprintf(&r, "%08b", s[i])
		}
	}
	return r.String()
}

// PrintUnicodeTable печатает таблицу.
func PrintUnicodeTable(begin, end rune) {
	for num := begin; num < end; num++ {
		symbol := string(num)
		dump := DumpRuneBinary(symbol)
		fmt.Printf("%5X  %s  %s\n", num, dump, symbol)
	}
}

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
		PrintUnicodeTable(rune(begin), rune(end))
	}
}
