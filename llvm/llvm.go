package llvm

import (
	"fmt"
	"strings"
)

func printComment(o *strings.Builder, comment string) {
	fmt.Fprintf(o, "; %s\n", comment)
}
