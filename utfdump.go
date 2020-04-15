package disgoft

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PrintBinaryDump Што То Делает
func PrintBinaryDump(filename string, f *os.File) {
	fmt.Print("Reading from ", filename, "\n")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			fmt.Print(DumpRuneBinary(string(c)), "\n")
		}
	}
}

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
