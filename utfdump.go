package disgoft

import (
	"fmt"
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
