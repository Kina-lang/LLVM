package llvm

import (
	"fmt"
	"strings"
)

func printComment(o *strings.Builder, comment string) {
	fmt.Fprintf(o, "; %s\n", comment)
}

func printBlankLine(o *strings.Builder) {
	fmt.Fprintf(o, "\n")
}

func indentLine(str string, level int) string {
	return fmt.Sprintf("%s%s", strings.Repeat("  ", level), str)
}
