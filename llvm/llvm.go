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
