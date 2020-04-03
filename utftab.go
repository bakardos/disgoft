package disgoft

import (
	"fmt"
)

// PrintUnicodeTable печатает таблицу.
func PrintUnicodeTable(begin, end rune) {
	for num := begin; num < end; num++ {
		symbol := string(num)
		dump := DumpRuneBinary(symbol)
		fmt.Printf("%5X  %s  %s\n", num, dump, symbol)
	}
}
